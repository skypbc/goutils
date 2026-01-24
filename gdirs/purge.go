package gdirs

import (
	"github.com/skypbc/goutils/gerrors"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

func Purge(path string, pattern string) error {
	if ok, err := Exists(path); !ok || err != nil {
		return nil
	}

	r, err := regexp.Compile(pattern)
	if err != nil {
		return gerrors.Wrap(err)
	}

	var errList []error
	if err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return gerrors.Wrap(err)
		}
		if !d.IsDir() && r.MatchString(path) {
			if err := os.Remove(path); err != nil {
				errList = append(errList, err)
			}
		}
		return nil
	}); err != nil {
		return gerrors.Wrap(err)
	}

	if len(errList) > 0 {
		params := make([]any, len(errList))
		for i, err := range errList {
			params[i] = err
		}
		return gerrors.NewUnknownError().
			SetTemplate("errors occurred during purge:\n{errors}").
			AddAny("errors", params)
	}

	return nil
}
