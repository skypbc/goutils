package utils

import (
	"github.com/skypbc/goutils/gbasic/gbreflect"
	"reflect"
)

func IsNil(v any) bool {
	if v == nil {
		return true
	}
	return gbreflect.IsNil(reflect.ValueOf(v))
}
