package gslice

func Intersect[S ~[]T, T comparable](s1 S, s2 S) (res S) {
	if len(s1) == 0 || len(s2) == 0 {
		return res
	}
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	res = make(S, 0, len(s1))
	if len(s1) == 0 {
		return res
	}
	u2 := map[T]struct{}{}
	for _, v := range s2 {
		u2[v] = struct{}{}
	}
	for _, v := range s1 {
		if _, ok := u2[v]; ok {
			res = append(res, v)
		}
	}
	return res
}

func Intersect2[S ~[]T, T any, K comparable](s1 S, s2 S, key func(v T) K) (res S) {
	if len(s1) == 0 || len(s2) == 0 {
		return res
	}
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	res = make(S, 0, len(s1))
	if len(s1) == 0 {
		return res
	}
	u2 := map[K]struct{}{}
	for _, v := range s2 {
		u2[key(v)] = struct{}{}
	}
	for _, v := range s1 {
		if _, ok := u2[key(v)]; ok {
			res = append(res, v)
		}
	}
	return res
}

func HasAny[S ~[]T, T comparable](s1 S, s2 S) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return false
	}
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	u2 := map[T]struct{}{}
	for _, v := range s2 {
		u2[v] = struct{}{}
	}
	for _, v := range s1 {
		if _, ok := u2[v]; ok {
			return true
		}
	}
	return false

}

func HasAny2[S ~[]T, T any, K comparable](s1 S, s2 S, key func(v T) K) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return false
	}
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	u2 := map[K]struct{}{}
	for _, v := range s2 {
		u2[key(v)] = struct{}{}
	}
	for _, v := range s1 {
		if _, ok := u2[key(v)]; ok {
			return true
		}
	}
	return false
}
