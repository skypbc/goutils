package gmap

import (
	"github.com/skypbc/goutils/gmap/internal"
	"reflect"
)

func Visit[M ~map[K]V, K comparable, V any](m M, f func(v any, keys []string)) {
	internal.Visit(reflect.ValueOf(m), nil, f)
}
