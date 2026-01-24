package internal

import (
	"github.com/skypbc/goutils/gfiles"
	"github.com/skypbc/goutils/gfmt/settings"
	"runtime"
)

type caller struct {
	File     string
	Line     int
	FuncName string
}

func GetCallerInfo(skip int) (res caller) {
	if pc, file, line, ok := runtime.Caller(skip + 1); ok {
		if details := runtime.FuncForPC(pc); details != nil {
			res.FuncName = details.Name()
		}
		res.File = gfiles.Normalize(file, settings.Print.Debug.PathSep)
		res.Line = line
	}
	return res
}
