package internal

import (
	"github.com/skypbc/goutils/gfmt"
	"reflect"
	"strconv"
)

func Visit(v reflect.Value, keys []string, f func(v any, keys []string)) {
	source := v

	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		child := make([]string, len(keys)+1)
		copy(child, keys)
		for i := 0; i < v.Len(); i++ {
			child[len(keys)] = strconv.Itoa(i)
			Visit(v.Index(i), child[:], f)
		}
	case reflect.Map:
		child := make([]string, len(keys)+1)
		copy(child, keys)
		for _, k := range v.MapKeys() {
			child[len(keys)] = gfmt.Sprintf("%v", k.Interface())
			Visit(v.MapIndex(k), child[:], f)
		}
	default:
		f(source.Interface(), keys)
	}
}

func Visit2(v reflect.Value, f func(v any)) {
	source := v

	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			Visit2(v.Index(i), f)
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			Visit2(v.MapIndex(k), f)
		}
	default:
		f(source.Interface())
	}
}

func Visit3(v reflect.Value, keys []string, f func(v any, keys []string)) {
	source := v

	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Map:
		child := make([]string, len(keys)+1)
		copy(child, keys)
		for _, k := range v.MapKeys() {
			child[len(keys)] = gfmt.Sprintf("%v", k.Interface())
			Visit3(v.MapIndex(k), child[:], f)
		}
	default:
		f(source.Interface(), keys)
	}
}
