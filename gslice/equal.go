package gslice

func Equal[S ~[]T, T comparable](s1 S, s2 S) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func EqualFunc[S ~[]T, T any](s1 S, s2 S, eq func(i int) bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if !eq(i) {
			return false
		}
	}
	return true
}
