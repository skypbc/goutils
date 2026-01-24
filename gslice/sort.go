package gslice

import (
	"github.com/skypbc/goutils/gslice/sortutils"
)

func Sort[S ~[]T, T any](slice S) {
	sortutils.Sort(slice)
}

func SortReverse[S ~[]T, T any](slice S) {
	sortutils.SortByFunc(slice, nil, true)
}

func SortByFunc[S ~[]T, T any](slice S, f func(i, j int) bool) {
	sortutils.SortByFunc(slice, f, false)
}

func SortByFuncReverse[S ~[]T, T any](slice S, f func(i, j int) bool) {
	sortutils.SortByFunc(slice, f, true)
}
