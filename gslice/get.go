package gslice

import (
	"github.com/skypbc/goutils/gerrors"
)

func Get[S ~[]T, T any](s S, index int) T {
	res, err := TryGet(s, index)
	if err != nil {
		panic(err)
	}
	return res
}

func TryGet[S ~[]T, T any](s S, index int) (T, error) {
	var zero T
	sLen := len(s)

	if index >= sLen || index < -sLen {
		return zero, gerrors.NewIndexOutOfRangeError()
	}

	if index < 0 {
		index = sLen + index
	}

	return s[index], nil
}

func Last[S ~[]T, T any](s S) (res T) {
	res, err := TryGet(s, -1)
	if err != nil {
		panic(err)
	}
	return res
}

func TryLast[S ~[]T, T any](s S) (res T, err error) {
	return TryGet(s, -1)
}
