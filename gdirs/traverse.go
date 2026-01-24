package gdirs

import (
	"github.com/rs/zerolog/log"
	"github.com/skypbc/goutils/gerrors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func Traverse(path string, deep int, fn func(entry fs.DirEntry, path string, deep int) error) (err error) {
	return Traverse2(nil, path, 0, deep, fn)
}

func Traverse2(root fs.DirEntry, rootPath string, deep int, maxDeep int, fn func(entry fs.DirEntry, path string, deep int) error) (err error) {
	if strings.HasSuffix(rootPath, ":") {
		rootPath += string(os.PathSeparator)
	}
	if (maxDeep >= 0 && maxDeep < deep) || fn == nil {
		return nil
	}
	if root == nil {
		fileInfo, err := os.Stat(rootPath)
		if err != nil {
			return gerrors.Wrap(err)
		}
		root = fs.FileInfoToDirEntry(fileInfo)
	}
	if err = fn(root, rootPath, deep); err != nil {
		return gerrors.Wrap(err)
	}
	if !root.IsDir() {
		return nil
	}
	items, err := os.ReadDir(rootPath)
	if err != nil {
		return gerrors.Wrap(err)
	}
	for _, entry := range items {
		path := filepath.Join(rootPath, entry.Name())
		if entry.IsDir() {
			if err = Traverse2(entry, path, deep+1, maxDeep, fn); err != nil {
				log.Error().Msg(gerrors.Wrap(err).Error())
			}
			continue
		}
		if err = fn(entry, path, deep); err != nil {
			return gerrors.Wrap(err)
		}
	}
	return nil
}
