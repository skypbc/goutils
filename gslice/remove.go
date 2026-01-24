package gslice

import (
	"github.com/skypbc/goutils/gerrors"
)

func Remove[S ~[]T, T any](s *S, index int) []T {
	sLen := len(*s)

	if index >= sLen || index < -sLen {
		panic(gerrors.NewIndexOutOfRangeError())
	}

	if index < 0 {
		index = sLen + index
	}

	if index == sLen-1 {
		*s = (*s)[:sLen-1]
		return *s
	}

	copy((*s)[:index], (*s)[index+1:])
	*s = (*s)[:sLen-1]

	return *s
}

func CloneAndRemove[S ~[]T, T any](s S, index int) []T {
	sLen := len(s)

	if index >= sLen || index < -sLen {
		panic(gerrors.NewIndexOutOfRangeError())
	}

	if index < 0 {
		index = sLen + index
	}

	res := make(S, len(s)-1)
	if index == sLen-1 {
		copy(res, s[:sLen-1])
		return res
	}

	copy(res, s[:index])
	copy(res[index:], s[index+1:])

	return res
}
