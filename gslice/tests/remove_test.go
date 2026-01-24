package tests

import (
	"github.com/skypbc/goutils/gerrors"
	"testing"
)

// go test -bench . -run notest -benchtime=10000x

var removeSource []int
var removeData []int

func init() {
	removeSource = make([]int, 1000000)
	removeData = make([]int, len(removeSource))

	for i := 0; i < len(removeData); i++ {
		removeSource[i] = i
	}
}

func removeDataReset() {
	copy(removeData, removeSource)
}

func Remove1[T any](s *[]T, index int) []T {
	sLen := len(*s)

	if index >= sLen || index < -sLen {
		panic(gerrors.NewIndexOutOfRangeError())
	}

	if index < 0 {
		index = sLen + index
	}

	for index+1 < sLen {
		(*s)[index] = (*s)[index+1]
		index++
	}

	*s = (*s)[:len(*s)-1]
	return *s
}

func Remove2[T any](s *[]T, index int) []T {
	sLen := len(*s)

	if index >= sLen || index < -sLen {
		panic(gerrors.NewIndexOutOfRangeError())
	}

	if index < 0 {
		index = sLen + index
	}

	*s = append((*s)[:index], (*s)[index+1:]...)[:sLen-1]
	return *s
}

func Remove3[T any](s *[]T, index int) []T {
	sLen := len(*s)

	if index >= sLen || index < -sLen {
		panic(gerrors.NewIndexOutOfRangeError())
	}

	if index < 0 {
		index = sLen + index
	}

	copy((*s)[:index], (*s)[index+1:])
	*s = (*s)[:sLen-1]

	return *s
}

func BenchmarkRemove1(b *testing.B) {
	removeDataReset()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if i > 0 {
			b.StartTimer()
		}

		Remove1(&removeData, 0)

		b.StopTimer()
		removeDataReset()
	}
}

func BenchmarkRemove2(b *testing.B) {
	removeDataReset()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if i > 0 {
			b.StartTimer()
		}

		Remove2(&removeData, 0)

		b.StopTimer()
		removeDataReset()
	}
}

func BenchmarkRemove3(b *testing.B) {
	removeDataReset()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if i > 0 {
			b.StartTimer()
		}

		Remove3(&removeData, 0)

		b.StopTimer()
		removeDataReset()
	}
}
