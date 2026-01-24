package gtime

import (
	"time"

	"golang.org/x/exp/constraints"
)

func Now() int64 {
	return time.Now().UTC().UnixNano()
}

func Second2Nano[T ~int64 | ~uint64](val T) T {
	return val * 1000000000
}

func Second2Nano2[T1 constraints.Integer, T2 ~int64 | ~uint64](val T1) T2 {
	return T2(val) * 1000000000
}

func Nano2Second[T ~int64 | ~uint64](val T) T {
	return val / 1000000000
}

func Nano2Second2[T1 constraints.Integer, T2 ~int64 | ~uint64](val T1) T2 {
	return T2(val) / 1000000000
}
