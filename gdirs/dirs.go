package gdirs

import (
	"errors"
	"github.com/skypbc/goutils/gerrors"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

func Create(path string, perm ...fs.FileMode) error {
	ok, err := Exists(path)
	if err != nil {
		return gerrors.Wrap(err)
	} else if ok {
		return nil
	}
	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		return gerrors.Wrap(err)
	}
	return nil
}

func Create2(path string, isFile bool, perm ...fs.FileMode) error {
	if isFile {
		path = filepath.Dir(path)
	}
	if err := Create(path); err != nil {
		return gerrors.Wrap(err)
	}
	return nil
}

func Delete(path string) error {
	if ok, err := Exists(path); err != nil {
		return gerrors.Wrap(err)
	} else if !ok {
		return nil
	}
	if err := os.RemoveAll(path); err != nil {
		return gerrors.Wrap(err)
	}
	return nil
}

func Exists(path string) (bool, error) {
	if info, err := os.Stat(path); err == nil {
		return info.IsDir(), nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, gerrors.Wrap(err)
	}
}

func Exists2(path string, isFile bool) (bool, error) {
	if isFile {
		path = filepath.Dir(path)
	}
	return Exists(path)
}

var (
	exeDirOnce sync.Once
	exeDirVal  string
	exeDirErr  error
)

// Executable возвращает абсолютный путь к директории, где лежит исполняемый файл. Вычисляется один раз и кешируется.
func Executable() (string, error) {
	exeDirOnce.Do(func() {
		exePath, err := os.Executable()
		if err != nil {
			exeDirErr = err
			return
		}
		// Нормализуем путь: раскрываем симлинки, чтобы получить "реальное" расположение бинарника
		exePath, err = filepath.EvalSymlinks(exePath)
		if err != nil {
			exeDirErr = err
			return
		}
		exeDirVal = filepath.Dir(exePath)
	})
	return exeDirVal, exeDirErr
}
