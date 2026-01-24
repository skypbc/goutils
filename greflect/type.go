package greflect

import (
	"reflect"
)

type kind uint8

const (
	None kind = iota
	Bool
	Int
	Uint
	Float
	String
	Complex
	Nil
	External
	Unknown
)

func Kind(v reflect.Value) kind {
	switch v.Kind() {
	case reflect.Bool:
		return Bool
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return Int
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return Uint
	case reflect.Float32, reflect.Float64:
		return Float
	case reflect.Complex64, reflect.Complex128:
		return Complex
	case reflect.String:
		return String
	default:
		if !v.IsValid() || v.IsZero() {
			return Nil
		}
	}
	return External
}

func Kind2(v reflect.Type) kind {
	switch v.Kind() {
	case reflect.Bool:
		return Bool
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return Int
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return Uint
	case reflect.Float32, reflect.Float64:
		return Float
	case reflect.Complex64, reflect.Complex128:
		return Complex
	case reflect.String:
		return String
	}
	return Unknown
}
