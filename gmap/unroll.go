package gmap

import (
	"github.com/skypbc/goutils/gmap/internal"
	"github.com/skypbc/goutils/gtypes"
	"reflect"
	"strings"
)

func Unroll[M ~map[K]V, K comparable, V any](m M) []gtypes.MapItem[string, any] {
	res := []gtypes.MapItem[string, any]{}

	internal.Visit(reflect.ValueOf(m), nil, func(v any, keys []string) {
		key := strings.Join(keys, ".")

		res = append(res, gtypes.MapItem[string, any]{
			Key:   key,
			Value: v,
		})
	})

	return res
}

func Unroll3[M ~map[K]V, K comparable, V any](m M) []gtypes.MapItem[string, any] {
	res := []gtypes.MapItem[string, any]{}

	internal.Visit3(reflect.ValueOf(m), nil, func(v any, keys []string) {
		key := strings.Join(keys, ".")

		res = append(res, gtypes.MapItem[string, any]{
			Key:   key,
			Value: v,
		})
	})

	return res
}
