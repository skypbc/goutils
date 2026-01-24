package gvars

import (
	"github.com/skypbc/goutils/gbasic/gbreflect"
	"reflect"
)

func IsMap(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Map
}

func IsSlice(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func IsArray(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Array
}

func Len(v any) int {
	return reflect.ValueOf(v).Len()
}

func IsNil(v any) bool {
	if v == nil {
		return true
	}
	return gbreflect.IsNil(reflect.ValueOf(v))
}

func IsNotNil(v any) bool {
	return !IsNil(v)
}
