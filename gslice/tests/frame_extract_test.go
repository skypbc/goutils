package tests

import (
	"fmt"
	"github.com/skypbc/goutils/gslice"
	"testing"
)

// go test -run TestExtract -v

func TestExtract1(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{2, 3}

	got := gslice.Extract(have, 1, 3)

	checkIntSlice(want, got, t)
}

func TestExtract2(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{1}

	got := gslice.Extract(have, 0, 1)

	checkIntSlice(want, got, t)
}

func TestExtract3(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{}

	got := gslice.Extract(have, 0, -1000)

	checkIntSlice(want, got, t)
}

func TestExtract4(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{}

	got := gslice.Extract(have, 0, -5)

	checkIntSlice(want, got, t)
}

func TestExtract5(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{1}

	got := gslice.Extract(have, 0, -4)

	checkIntSlice(want, got, t)
}

func TestExtract6(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{}

	got := gslice.Extract(have, 1000, -1)

	checkIntSlice(want, got, t)
}

func TestExtract7(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{1, 2, 3, 4}

	got := gslice.Extract(have, -5, -1)

	checkIntSlice(want, got, t)
}

func TestExtract8(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{2, 3, 4}

	got := gslice.Extract(have, -4, -1)

	checkIntSlice(want, got, t)
}

func TestExtract9(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{1, 2, 3, 4}

	got := gslice.Extract(have, -1000, -1)

	checkIntSlice(want, got, t)
}

func TestExtract10(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{1, 2, 3, 4, 5}

	got := gslice.Extract(have, 0, 1000)

	checkIntSlice(want, got, t)
}

func TestExtract11(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{1, 2, 3, 4, 5}

	got := gslice.Extract(have, -1000, 1000)

	checkIntSlice(want, got, t)
}

func TestExtract12(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	got := gslice.Extract(have, 0, 3)

	got[0] = 1000

	if have[0] == 1000 {
		t.Error(fmt.Errorf("want got[0] != have[0], got[%d] != have[%d]...", got[0], have[0]))
	}
}
