package gslice

import (
	"github.com/skypbc/goutils/gerrors"
)

func Pop[S ~[]T, T any](s *S, index int) T {
	res, err := TryPop(s, index)
	if err != nil {
		panic(err)
	}
	return res
}

func TryPop[S ~[]T, T any](s *S, index int) (T, error) {
	var res T
	sLen := len(*s)

	if index >= sLen || index < -sLen {
		return res, gerrors.NewIndexOutOfRangeError()
	}

	if index < 0 {
		index = sLen + index
	}

	res = (*s)[index]

	if index == sLen-1 {
		*s = (*s)[:sLen-1]
		return res, nil
	}

	copy((*s)[:index], (*s)[index+1:])
	*s = (*s)[:sLen-1]

	return res, nil
}
