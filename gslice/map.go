package gslice

func Map[S ~[]T, T any](s S, f func(i int) T) S {
	res := make(S, len(s))
	for i := range s {
		res[i] = f(i)
	}
	return res
}

func Map2[S ~[]T, T any](s S, f func(v T) T) S {
	res := make(S, len(s))
	for i, v := range s {
		res[i] = f(v)
	}
	return res
}
