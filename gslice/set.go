package gslice

import (
	"github.com/skypbc/goutils/gerrors"
)

func Set[S ~[]T, T any](s S, index int, value T) {
	err := TrySet(s, index, value)
	if err != nil {
		panic(err)
	}
}

func TrySet[S ~[]T, T any](s S, index int, value T) error {
	sLen := len(s)

	if index >= sLen || index < -sLen {
		return gerrors.NewIndexOutOfRangeError()
	}

	if index < 0 {
		index = sLen + index
	}

	s[index] = value
	return nil
}

func SetAndGet[S ~[]T, T any](s S, index int, value T) T {
	res, err := TrySetAndGet(s, index, value)
	if err != nil {
		panic(err)
	}
	return res
}

func TrySetAndGet[S ~[]T, T any](s S, index int, value T) (old T, res error) {
	var zero T
	sLen := len(s)

	if index >= sLen || index < -sLen {
		return zero, gerrors.NewIndexOutOfRangeError()
	}

	if index < 0 {
		index = sLen + index
	}

	old = s[index]
	s[index] = value

	return
}
