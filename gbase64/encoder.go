package gbase64

import (
	"encoding/base64"
	"io"
)

// NewEncoder создает потоковый энкодер для кодирования данных в base64.
func NewEncoder(enc *base64.Encoding, in io.Reader, minSize ...int) (out io.ReadCloser) {
	// 384 = 512 / 4 * 3
	size := 384
	if len(minSize) > 0 && minSize[0] > 0 {
		size = minSize[0]
	}
	// Получившееся число, должно делиться ровно на 4.
	// Округляется в большу сторону.
	size = (size + (3 - (size % 3))) / 3 * 4

	s := &streamEncoder{
		buff:   make([]byte, size),
		source: in,
	}
	s.encoder = base64.NewEncoder(enc, s)

	return s
}

type streamEncoder struct {
	source io.Reader

	buff   []byte
	offset int
	size   int

	encoder io.WriteCloser
	closed  bool
}

// Close закрывает энкодер и освобождает ресурсы. После закрытия, необходимо вызвать Read для получения оставшейся
// части данных, которые могут возниктуть из-за специфики кодирования в base64. Как правило, это 1-2 символа.
func (s *streamEncoder) Close() error {
	s.closed = true
	return s.encoder.Close()
}

// Write используется внутренним энкодером для записи закодированных base64 данных во временный буфер.
func (s *streamEncoder) Write(p []byte) (n int, err error) {
	// Такое не должно произойти
	if size := len(p); len(s.buff[s.size:]) < size {
		return 0, io.ErrShortBuffer
	}
	// Добавляем новые данные в конец
	n += copy(s.buff[s.size:], p)
	s.size += n
	return n, nil
}

// Read возвращает закодированные данные.
func (s *streamEncoder) Read(p []byte) (n int, err error) {
	for {
		// В буфере есть данные для передачи?
		if s.offset < s.size {
			size := copy(p, s.buff[s.offset:s.size])
			n += size
			s.offset += size
			p = p[size:]
			// Все данные были скопированы?
			if s.offset == s.size {
				s.offset = 0
				s.size = 0
			}
			// Мы заполнили входящий буфер?
			if len(p) == 0 {
				return n, nil
			}
		}
		// Определяем максимальный размер сырых данных, который может поместиться в буфер после кодирования в base64.
		// Округляем в меньшую сторону.
		maxRaw := len(s.buff[s.offset:]) / 4 * 3
		if len(p) < maxRaw {
			maxRaw = len(p)
		}

		readed := 0
		// Мы должны пропустить шаг, если получим ошибку (io.EOF) с прошлого круга.
		// Она сигнализирует о завершения чтения.
		if err == nil {
			// Читаем новую порцию данных
			readed, err = s.source.Read(p[:maxRaw])
		}

		if err != nil {
			// Если дочитали до конца
			if err == io.EOF {
				// И не получили данных
				if readed == 0 {
					// Мы отдали все данные
					if s.closed {
						return n, err
					}
					// После закрытия, могут появиться данные
					if err2 := s.Close(); err2 != nil {
						// Неизвестная ошибка
						return n, err2
					}
					// Возвращаемся дочитывать данные
					continue
				}
			} else {
				return n, err
			}
		}
		// s.encoder.Write(p[:readed]) //nolint:errcheck
		if _, err := s.encoder.Write(p[:readed]); err != nil {
			// Неизвестная ошибка
			return n, err
		}
	}
}
