package tests

import (
	"github.com/skypbc/goutils/gslice"
	"testing"
)

// go test -run TestAppend -v
// go test -bench . -run notest -benchtime=10000x

func TestAppend1(t *testing.T) {
	have := []int{1, 2, 3, 4, 5, 1111}
	want := []int{1, 2, 3, 4, 5, 1000}

	got := have[:5]
	gslice.Append(&got, 1000)

	checkIntSlice(have, got, t)
	checkIntSlice(want, got, t)
}
