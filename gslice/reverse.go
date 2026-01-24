package gslice

func Reverse[S ~[]T, T comparable](s S) S {
	len := len(s)
	mid := len / 2

	for i := 0; i < mid; i++ {
		j := len - i - 1

		s[i], s[j] = s[j], s[i]
	}
	return s
}
