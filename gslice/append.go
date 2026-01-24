package gslice

func Append[S ~[]T, T any](s *S, values ...T) S {
	*s = append(*s, values...)
	return *s
}
