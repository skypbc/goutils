package gdirs

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type CopyOpts struct {
	OnCopy  func(root string, src string, dst string, srcInfo fs.FileInfo) (skip bool, err error)
	OnError func(root string, src string, dst string, err error) (errOut error)
}

// Copy рекурсивно копирует директорию src в dst. Права файлов и директорий сохраняются.
// Симлинки копируются как симлинки.
func Copy(src string, dst string, opts ...CopyOpts) (err error) {
	if src, err = filepath.Abs(src); err != nil {
		return err
	}
	if dst, err = filepath.Abs(dst); err != nil {
		return err
	}

	if src == dst {
		return &fs.PathError{
			Op:   "copy",
			Path: src,
			Err:  fs.ErrInvalid,
		}
	}

	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !srcInfo.IsDir() {
		return &fs.PathError{
			Op:   "copy",
			Path: src,
			Err:  fs.ErrInvalid,
		}
	}

	// Создаем корневую директорию назначения
	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		defer func() {
			if err != nil && len(opts) > 0 && opts[0].OnError != nil {
				err = opts[0].OnError(src, path, dst, err)
			}
		}()

		if err != nil {
			return err
		}

		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)

		info, err := d.Info()
		if err != nil {
			return err
		}

		if len(opts) > 0 && opts[0].OnCopy != nil {
			skip, err := opts[0].OnCopy(src, path, target, info)
			if skip {
				return nil
			}
			if err != nil {
				return err
			}
		}

		switch {
		case d.IsDir():
			if rel == "." {
				return nil
			}
			return os.MkdirAll(target, info.Mode())

		case info.Mode()&os.ModeSymlink != 0:
			link, err := os.Readlink(path)
			if err != nil {
				return err
			}
			return os.Symlink(link, target)

		default:
			return copyFile(path, target, info.Mode())
		}
	})
}

func copyFile(src, dst string, perm fs.FileMode) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	return out.Sync()
}
