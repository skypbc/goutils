package gslice

func Convert[S ~[]T, T any, R any](s S, f func(i int) R) []R {
	res := make([]R, len(s))
	for i := range s {
		res[i] = f(i)
	}
	return res
}

func Convert2[S ~[]T, T any, R any](s S, f func(v T) R) []R {
	res := make([]R, len(s))
	for i, v := range s {
		res[i] = f(v)
	}
	return res
}

func ConvertAndFilter[S ~[]T, T any, R any](s S, f func(i int) (R, bool)) []R {
	res := make([]R, 0, len(s))
	for i := range s {
		if r, include := f(i); include {
			res = append(res, r)
		}
	}
	return res
}

func ConvertAndFilter2[S ~[]T, T any, R any](s S, f func(v T) (R, bool)) []R {
	res := make([]R, 0, len(s))
	for _, v := range s {
		if r, include := f(v); include {
			res = append(res, r)
		}
	}
	return res
}
