package gslice

func Visit[S ~[]T, T any](s S, f func(i int)) {
	for i := range s {
		f(i)
	}
}

func Visit2[S ~[]T, T any](s S, f func(i int, v T)) {
	for i, v := range s {
		f(i, v)
	}
}
