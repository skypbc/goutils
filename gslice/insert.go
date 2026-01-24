package gslice

import (
	"github.com/skypbc/goutils/gerrors"
	"github.com/skypbc/goutils/gnum"
)

func Insert[S ~[]T, T any](s *S, index int, val T) S {
	res, err := TryInsert(s, index, val, true)
	if err != nil {
		panic(err)
	}
	return res
}

func Insert2[S ~[]T, T any](s *S, index int, val T, before bool) S {
	res, err := TryInsert(s, index, val, before)
	if err != nil {
		panic(err)
	}
	return res
}

func TryInsert[S ~[]T, T any](s *S, index int, val T, before bool) (S, error) {
	sLen := len(*s)

	if index < 0 {
		index = sLen + index
	}

	if !before {
		index++
	}

	if sLen == 0 && index == -1 {
		index = 0
	}

	if before && (index > sLen || index < -sLen) {
		return nil, gerrors.NewIndexOutOfRangeError()
	}

	if index == sLen {
		*s = append(*s, val)
		return *s, nil
	}

	pre := (*s)[:index]
	post := (*s)[index:]

	nLen := gnum.Max(index+1, sLen+1)
	sCap := cap(*s)

	if nLen > sCap {
		if sCap < 1024 {
			*s = make([]T, nLen, sCap+sCap)
		} else {
			*s = make([]T, nLen, int(float32(sCap)*1.25))
		}
	} else {
		*s = make([]T, nLen, sCap)
	}

	n := copy(*s, pre)

	(*s)[n] = val
	n++

	copy((*s)[n:], post)

	return *s, nil
}
