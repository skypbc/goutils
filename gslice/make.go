package gslice

func Make[S ~[]T, T any](s S, args ...int) (res S) {
	switch len(args) {
	case 1:
		res = make(S, args[0])
	case 2:
		res = make(S, args[0], args[1])
	default:
		res = make(S, 0)
	}
	return res
}
