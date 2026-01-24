package gslice

func Subtract[S ~[]T, T comparable](src, sub S) S {
	result := make(S, 0, len(src))
	if len(sub) == 0 {
		result = append(result, src...)
		return result
	}
	remove := make(map[T]struct{}, len(sub))
	for _, v := range sub {
		remove[v] = struct{}{}
	}
	for _, v := range src {
		if _, found := remove[v]; !found {
			result = append(result, v)
		}
	}
	return result
}

func SubtractFunc[S ~[]T, T any, K comparable](src, sub S, f func(T) K) S {
	result := make(S, 0, len(src))
	if len(sub) == 0 {
		result = append(result, src...)
		return result
	}
	remove := make(map[K]struct{}, len(sub))
	for _, v := range sub {
		remove[f(v)] = struct{}{}
	}
	for _, v := range src {
		if _, found := remove[f(v)]; !found {
			result = append(result, v)
		}
	}
	return result
}
