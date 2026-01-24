package gslice

import (
	"math/rand"
)

func Shuffle[S ~[]T, T any](s S) {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

func Shuffle2[S ~[]T, T any](s S) S {
	tmp := make(S, len(s))
	copy(tmp, s)
	Shuffle(tmp)
	return tmp
}

func Choice[S ~[]T, T any](s S) T {
	return s[rand.Intn(len(s))]
}
