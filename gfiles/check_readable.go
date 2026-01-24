package gfiles

import (
	"context"
	"errors"
	"fmt"
	"github.com/skypbc/goutils/gerrors"
	"io"
	"os"
	"time"
)

// ErrInsufficientReadable возвращается, если удалось прочитать меньше, чем требовалось.
type ErrInsufficientReadable struct {
	Required int64
	Read     int64
	Reason   error // первичная причина (например, EOF, sharing violation и т.п.)
}

func (e *ErrInsufficientReadable) Error() string {
	return fmt.Sprintf("readable bytes %d < required %d: %v", e.Read, e.Required, e.Reason)
}

func (e *ErrInsufficientReadable) Unwrap() error { return e.Reason }

// WaitReadableN проверяет, что из файла можно прочитать как минимум requiredBytes,
// читая его маленькими блоками (chunkSize), чтобы не расходовать память.
// Если requiredBytes < 0 — означает "весь текущий файл".
// Если chunkSize <= 0 — используется значение по умолчанию (64 КБ).
//
// Функция делает ретраи открытия/первого чтения до дедлайна контекста (актуально для Windows,
// где возможны кратковременные блокировки). После успешного открытия чтение происходит
// в одном проходе с поддержкой отмены по контексту.
//
// Важно: если файл короче requiredBytes, вернётся ErrInsufficientReadable.
func WaitReadableN(
	ctx context.Context, path string, requiredBytes int64, chunkSize int, timeout ...time.Duration,
) error {
	// Контекст с дедлайном (если указан таймаут), если нет, то ожидается что ctx уже с дедлайном
	if len(timeout) > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout[0])
		defer cancel()
	}

	if chunkSize <= 0 {
		chunkSize = 64 * 1024 // 64 KiB
	} else if chunkSize < 512 {
		// Слишком маленькие блоки могут сильно замедлять I/O
		chunkSize = 512
	}

	// Первичная проверка: существует и это обычный файл
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !info.Mode().IsRegular() {
		return errors.New("path is not a regular file")
	}

	// Рассчитываем целевой объём к проверке
	targetSize := requiredBytes
	if requiredBytes < 0 {
		targetSize = info.Size() // "весь файл на момент проверки"
	}
	if targetSize == 0 {
		// Пустой файл: достаточно проверить, что его можно открыть
		// (встречаются случаи эксклюзивных блокировок без чтения).
		f, err := tryOpenWithRetries(ctx, path)
		if err != nil {
			return err
		}
		f.Close()
		return nil
	}

	// Открываем с ретраями (на случай sharing violation)
	f, err := tryOpenWithRetries(ctx, path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Читаем кусками до достижения target или конца файла
	buf := make([]byte, chunkSize)
	var total int64

	for total < targetSize {
		// Контекстная отмена между чтениями
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		toRead := chunkSize
		remaining := targetSize - total
		if remaining < int64(toRead) {
			toRead = int(remaining)
		}

		n, rerr := f.Read(buf[:toRead])
		if n > 0 {
			total += int64(n)
		}
		if rerr != nil {
			// EOF — это не "ошибка" сама по себе, но означает, что дальше читать нечего.
			if errors.Is(rerr, io.EOF) {
				if total >= targetSize {
					return nil
				}
				return &ErrInsufficientReadable{
					Required: targetSize,
					Read:     total,
					Reason:   rerr,
				}
			}
			// Любая иная ошибка чтения
			return &ErrInsufficientReadable{
				Required: targetSize,
				Read:     total,
				Reason:   rerr,
			}
		}
	}

	return nil
}

// IsReadableOnceN — упрощённый вариант без ретраев.
// Возвращает (ok, error). ok=true, если прочитан нужный объём.
func IsReadableOnceN(path string, requiredBytes int64, chunkSize int) (bool, error) {
	// Контекст без дедлайна: одна попытка (без повторов открытия)
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if !info.Mode().IsRegular() {
		return false, errors.New("path is not a regular file")
	}
	target := requiredBytes
	if requiredBytes < 0 {
		target = info.Size()
	}
	if target == 0 {
		f, err := os.Open(path)
		if err != nil {
			return false, err
		}
		f.Close()
		return true, nil
	}

	if chunkSize <= 0 {
		chunkSize = 64 * 1024
	} else if chunkSize < 512 {
		chunkSize = 512
	}

	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	buf := make([]byte, chunkSize)
	var total int64
	for total < target {
		toRead := chunkSize
		remaining := target - total
		if remaining < int64(toRead) {
			toRead = int(remaining)
		}
		n, rerr := f.Read(buf[:toRead])
		if n > 0 {
			total += int64(n)
		}
		if rerr != nil {
			if errors.Is(rerr, io.EOF) {
				return total >= target, nil
			}
			return false, rerr
		}
	}
	return true, nil
}

// tryOpenWithRetries — открытие файла с небольшими ретраями до дедлайна контекста.
// Полезно на Windows, где возможны ERROR_SHARING_VIOLATION при коротких гонках.
func tryOpenWithRetries(ctx context.Context, path string) (*os.File, error) {
	const (
		initialBackoff = 30 * time.Millisecond
		maxBackoff     = 300 * time.Millisecond
	)
	backoff := initialBackoff

	var lastErr error
	for {
		f, err := os.Open(path)
		if err == nil {
			return f, nil
		}
		lastErr = err

		select {
		case <-ctx.Done():
			if cerr := ctx.Err(); cerr != nil {
				lastErr = gerrors.Wrap(lastErr, cerr).
					SetTemplate(`context done while trying to open file "{file}"`).
					AddStr("file", path)
			}
			return nil, lastErr
		default:
		}

		// Если нет дедлайна — делаем одну попытку
		if _, hasDeadline := ctx.Deadline(); !hasDeadline {
			return nil, lastErr
		}

		// Подождём и попробуем снова
		sleep := backoff
		backoff *= 2
		if backoff > maxBackoff {
			backoff = maxBackoff
		}
		timer := time.NewTimer(sleep)
		select {
		case <-ctx.Done():
			timer.Stop()
			return nil, ctx.Err()
		case <-timer.C:
		}
	}
}
