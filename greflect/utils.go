package greflect

import (
	"github.com/skypbc/goutils/gbasic/gbreflect"
	"reflect"
)

var IsNil = gbreflect.IsNil
var IsPointer = gbreflect.IsPointer

func Elem(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer,
		reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		if v.IsNil() {
			return v
		}
	}
	if v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
		return v.Elem()
	}
	return v
}

func TryElem(v reflect.Value) (reflect.Value, bool) {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer,
		reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		if v.IsNil() {
			return v, false
		}
	}
	if v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
		return v.Elem(), true
	}
	return v, false
}

func LastElem(v reflect.Value) reflect.Value {
	for {
		switch v.Kind() {
		case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer,
			reflect.UnsafePointer, reflect.Interface, reflect.Slice:
			if v.IsNil() {
				return v
			}
		}
		if v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
			v = v.Elem()
		} else {
			break
		}
	}
	return v
}
