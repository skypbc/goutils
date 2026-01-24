package gnum

import (
	cs "golang.org/x/exp/constraints"
)

func Max[T cs.Integer | cs.Float](v1 T, v2 T) T {
	if v1 > v2 {
		return v1
	}
	return v2
}

func Min[T cs.Integer | cs.Float](v1 T, v2 T) T {
	if v1 < v2 {
		return v1
	}
	return v2
}
