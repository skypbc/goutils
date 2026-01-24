package gset

func FromSlice[S ~[]T, T comparable](s S) (res map[T]struct{}) {
	res = make(map[T]struct{}, len(s))
	for _, v := range s {
		res[v] = struct{}{}
	}
	return res
}

func FromSliceFunc[S ~[]T1, T1 any, T2 comparable](s S, f func(i int) T2) (res map[T2]struct{}) {
	res = make(map[T2]struct{}, len(s))
	for i := range s {
		res[f(i)] = struct{}{}
	}
	return res
}

func FromMapKeys[M ~map[K]V, K comparable, V any](m M) (res map[K]struct{}) {
	res = make(map[K]struct{}, len(m))
	for k := range m {
		res[k] = struct{}{}
	}
	return res
}

func FromMapValues[M ~map[K]V, K comparable, V comparable](m M) (res map[V]struct{}) {
	res = make(map[V]struct{}, len(m))
	for _, v := range m {
		res[v] = struct{}{}
	}
	return res
}

func FromMapValuesFunc[M ~map[K]V, K comparable, V any, T comparable](m M, f func(v V) T) (res map[T]struct{}) {
	res = make(map[T]struct{}, len(m))
	for _, v := range m {
		res[f(v)] = struct{}{}
	}
	return res
}
