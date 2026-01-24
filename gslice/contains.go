package gslice

import "github.com/skypbc/goutils/gslice/utils"

func Contains[S ~[]T, T comparable](s S, v T) bool {
	return Index(s, v) >= 0
}

func ContainsFunc[S ~[]T, T any](s S, f func(i int) bool) bool {
	return IndexFunc(s, f) >= 0
}

func ContainsNil[S ~[]T, T any](s S) bool {
	return IndexFunc(s, func(i int) bool {
		return utils.IsNil(s[i])
	}) != -1
}
