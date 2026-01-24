package gdebug

import (
	"github.com/skypbc/goutils/gslice"
	"reflect"
	"runtime"
	"strings"
)

func FunctionName(f any) string {
	fname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return gslice.Get(strings.Split(fname, "."), -1)
}
