package gslice

func Extract[S ~[]T, T any](s S, from int, to int) S {
	var res []T

	view := View(s, from, to)
	if len(view) == 0 {
		return res
	}

	res = make([]T, len(view))
	for i, val := range view {
		res[i] = val
	}
	return res
}

func View[S ~[]T, T any](s S, from int, to int) S {
	sLen := len(s)
	if from >= sLen || to < -sLen {
		return []T{}
	}

	for from < 0 {
		from = sLen + from
	}

	if to > sLen {
		to = sLen
	}

	if to < 0 {
		to = sLen + to
	}

	return s[from:to]
}
