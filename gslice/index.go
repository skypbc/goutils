package gslice

func Index[S ~[]T, T comparable](s S, v T) int {
	for i, v_ := range s {
		if v == v_ {
			return i
		}
	}
	return -1
}

func IndexFunc[S ~[]T, T any](s S, f func(i int) bool) int {
	for i := range s {
		if f(i) {
			return i
		}
	}
	return -1
}
