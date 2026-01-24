// Contains functions for sorting generic map[K]V by keys or by values, or using function. It returns []MapItem[K, V].
package gmap

import (
	"github.com/skypbc/goutils/gmap/sortutils"
	"github.com/skypbc/goutils/gtypes"
)

// Sort function sorts a map by keys
func Sort[M ~map[K]V, K comparable, V any](
	m M,
) []gtypes.MapItem[K, V] {
	return sortutils.SortByFunc(m, nil, false, false)
}

// SortReverse function sort a map by keys (reverse order)
func SortReverse[M ~map[K]V, K comparable, V any](
	m M,
) []gtypes.MapItem[K, V] {
	return sortutils.SortByFunc(m, nil, true, false)
}

// SortByKey function sorts a map by keys
func SortByKey[M ~map[K]V, K comparable, V any](
	m M,
) []gtypes.MapItem[K, V] {
	return sortutils.SortByFunc(m, nil, false, false)
}

// SortByKeyReverse function sorts a map by keys (reverse order)
func SortByKeyReverse[M ~map[K]V, K comparable, V any](
	m M,
) []gtypes.MapItem[K, V] {
	return sortutils.SortByFunc(m, nil, true, false)
}

// SortByValue function sorts a map by values
func SortByValue[M ~map[K]V, K comparable, V any](
	m M,
) []gtypes.MapItem[K, V] {
	return sortutils.SortByFunc(m, nil, false, true)
}

// SortByValue function sorts a map by values (reverse order)
func SortByValueReverse[M ~map[K]V, K comparable, V any](
	m M,
) []gtypes.MapItem[K, V] {
	return sortutils.SortByFunc(m, nil, true, true)
}

// SortByFunc function sorts a map using a function
func SortByFunc[M ~map[K]V, K comparable, V any](
	m M,
	f func(p1, p2 gtypes.MapItem[K, V]) bool,
) []gtypes.MapItem[K, V] {
	return sortutils.SortByFunc(m, f, false, false)
}

// SortByFuncReverse function sorts a map using a function (reverse order)
func SortByFuncReverse[M ~map[K]V, K comparable, V any](
	m M,
	f func(p1, p2 gtypes.MapItem[K, V]) bool,
) []gtypes.MapItem[K, V] {
	return sortutils.SortByFunc(m, f, true, false)
}
