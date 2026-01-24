package gbreflect

import "reflect"

func IsNil(v reflect.Value) bool {
	return !v.IsValid() || (IsPointer(v) && v.IsNil())
}

func IsPointer(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer:
		return true
	case reflect.Interface, reflect.Slice:
		return true
	}
	return false
}
