package gslice

func Unique[S ~[]T, T comparable](s S) (res S) {
	if len(s) <= 1 {
		return s
	}
	seen := make(map[T]struct{}, len(s))
	cursor := 0
	for index, v := range s {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			s[cursor] = s[index]
			cursor++
		}
	}
	return s[:cursor]
}

func UniqueFunc[S ~[]T, T any, K comparable](s S, f func(i int) K) (res S) {
	if len(s) <= 1 {
		return s
	}
	seen := make(map[K]struct{}, len(s))
	cursor := 0
	for index := range s {
		k := f(index)
		if _, ok := seen[k]; !ok {
			seen[k] = struct{}{}
			s[cursor] = s[index]
			cursor++
		}
	}
	return s[:cursor]
}
