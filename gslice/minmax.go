package gslice

import (
	"cmp"
	"github.com/skypbc/goutils/gerrors"
)

func Min[S ~[]T, T cmp.Ordered](s S) T {
	if len(s) < 1 {
		err := gerrors.NewSliceError().
			SetTemplate("slice is empty")
		panic(err)
	}
	m := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] < m {
			m = s[i]
		}
	}
	return m
}

func MinFunc[S ~[]T, T any](s S, less func(i, j int) bool) T {
	if len(s) < 1 {
		err := gerrors.NewSliceError().
			SetTemplate("slice is empty")
		panic(err)
	}
	i := 0
	for j := 1; j < len(s); j++ {
		if !less(i, j) {
			i = j
		}
	}
	return s[i]
}

func Max[S ~[]T, T cmp.Ordered](s S) T {
	if len(s) < 1 {
		err := gerrors.NewSliceError().
			SetTemplate("slice is empty")
		panic(err)
	}
	m := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > m {
			m = s[i]
		}
	}
	return m
}

func MaxFunc[S ~[]T, T any](s S, more func(i, j int) bool) T {
	if len(s) < 1 {
		err := gerrors.NewSliceError().
			SetTemplate("slice is empty")
		panic(err)
	}
	i := 0
	for j := 1; j < len(s); j++ {
		if !more(i, j) {
			i = j
		}
	}
	return s[i]
}
