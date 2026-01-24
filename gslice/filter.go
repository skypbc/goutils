package gslice

func Filter[S ~[]T, T any](s S, f func(i int) (include bool)) (res S) {
	ids := make([]int, 0, len(s))
	for i := range s {
		if f(i) {
			ids = append(ids, i)
		}
	}
	res = make([]T, 0, len(ids))
	for _, index := range ids {
		res = append(res, s[index])
	}
	return res
}

func Filter2[S ~[]T, T any](s S, f func(v T) (include bool)) (res S) {
	res = make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}
