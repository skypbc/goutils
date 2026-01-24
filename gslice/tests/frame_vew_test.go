package tests

import (
	"github.com/skypbc/goutils/gslice"
	"testing"
)

// go test -run TestView -v

func TestView1(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{2, 3}

	got := gslice.View(have, 1, 3)

	checkIntSlice(want, got, t)
}

func TestView2(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{1}

	got := gslice.View(have, 0, 1)

	checkIntSlice(want, got, t)
}

func TestView3(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{}

	got := gslice.View(have, 0, -1000)

	checkIntSlice(want, got, t)
}

func TestView4(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{}

	got := gslice.View(have, 0, -5)

	checkIntSlice(want, got, t)
}

func TestView5(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{1}

	got := gslice.View(have, 0, -4)

	checkIntSlice(want, got, t)
}

func TestView6(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{}

	got := gslice.View(have, 1000, -1)

	checkIntSlice(want, got, t)
}

func TestView7(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{1, 2, 3, 4}

	got := gslice.View(have, -5, -1)

	checkIntSlice(want, got, t)
}

func TestView8(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{2, 3, 4}

	got := gslice.View(have, -4, -1)

	checkIntSlice(want, got, t)
}

func TestView9(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{1, 2, 3, 4}

	got := gslice.View(have, -1000, -1)

	checkIntSlice(want, got, t)
}

func TestView10(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{1, 2, 3, 4, 5}

	got := gslice.View(have, 0, 1000)

	checkIntSlice(want, got, t)
}

func TestView11(t *testing.T) {
	have := []int{1, 2, 3, 4, 5}
	want := []int{1, 2, 3, 4, 5}

	got := gslice.View(have, -1000, 1000)

	checkIntSlice(want, got, t)
}
