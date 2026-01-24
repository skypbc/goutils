package gslice

import "github.com/skypbc/goutils/gnum"

func Clone[S ~[]T, T any](s S) S {
	res := make(S, len(s), cap(s))
	copy(res, s)
	return res
}

func CloneTight[S ~[]T, T any](s S) S {
	res := make(S, len(s))
	copy(res, s)
	return res
}

func CloneAndResize[S ~[]T, T any](s S, size int, capacity int) S {
	res := make(S, size, gnum.Max(size|len(s), capacity|cap(s)))
	copy(res, s)
	return res
}

func CloneAndAppend[S ~[]T, T any](s S, val ...T) S {
	res := make(S, len(s)+len(val))
	copy(res, s)
	for i := 0; i < len(val); i++ {
		res[len(s)+i] = val[i]
	}
	return res
}

func CloneTo[S ~[]T, T any](s S, out *S) S {
	*out = make([]T, len(s), cap(s))
	copy(*out, s)
	return *out
}

func CloneAndResizeTo[S ~[]T, T any](s S, out *S, size int, capacity int) S {
	*out = make([]T, size, gnum.Max(size|len(s), capacity|cap(s)))
	copy(*out, s)
	return *out
}
