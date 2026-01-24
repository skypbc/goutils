package gmap

func Make[M ~map[K]V, K comparable, V any](m M, args ...int) M {
	if len(args) > 0 {
		return make(M, args[0])
	}
	return make(M, 0)
}
