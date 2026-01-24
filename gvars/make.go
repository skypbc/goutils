package gvars

func Make[T any](v T) (res T) {
	return res
}

// fs = format slice
func Makefs[S ~[]T, T any](s S) (res T) {
	return res
}

// fmk = format map, res is key type
func Makefmk[M ~map[K]V, K comparable, V any](m M) (res K) {
	return res
}

// fmv = format map, res is value type
func Makefmv[M ~map[K]V, K comparable, V any](m M) (res V) {
	return res
}
