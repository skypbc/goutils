package tests

import (
	"fmt"
	"github.com/skypbc/goutils/gslice"
	"testing"
)

// go test -run TestInsert -v

func TestInsertToStart1(t *testing.T) {
	have := getTestData()
	want := []int{
		1000, 1, 2, 3, 4, 5,
	}

	got := have[:len(have):len(have)]
	gslice.Insert(&got, 0, 1000)

	checkIntSlice(want, got, t)
}

func TestInsertToStart2(t *testing.T) {
	have := getTestData()
	want := []int{
		1000, 1, 2, 3, 4, 5,
	}

	got := have[:len(have):len(have)]
	gslice.Insert(&got, -5, 1000)

	checkIntSlice(want, got, t)
}

func TestInsertToStartWithError1(t *testing.T) {
	have := getTestData()
	want := []int{
		1, 2, 3, 4, 5,
	}

	got := have[:len(have):len(have)]
	_, err := gslice.TryInsert(&got, -6, 1000, true)

	checkIntSlice(want, got, t)

	if err == nil {
		t.Error(fmt.Errorf("want err != nil, got err == nil..."))
	}
}

func TestInsertToStartWithError2(t *testing.T) {
	have := getTestData()
	want := []int{
		1, 2, 3, 4, 5,
	}
	got := have[:len(have):len(have)]
	_, err := gslice.TryInsert(&got, 6, 1000, true)

	checkIntSlice(want, got, t)

	if err == nil {
		t.Error(fmt.Errorf("want err != nil, got err == nil..."))
	}
}

func TestInsertToEnd(t *testing.T) {
	have := getTestData()
	want := []int{
		1, 2, 3, 4, 5, 1000,
	}

	got := have[:len(have):len(have)]
	gslice.Insert(&got, len(got), 1000)

	checkIntSlice(want, got, t)
}

func TestInsertToMinus1(t *testing.T) {
	have := getTestData()
	want := []int{
		1, 2, 3, 4, 1000, 5,
	}

	got := have[:len(have):len(have)]
	gslice.Insert(&got, -1, 1000)

	checkIntSlice(want, got, t)
}

func TestInsertToMinus2(t *testing.T) {
	have := getTestData()
	want := []int{
		1, 2, 3, 1000, 4, 5,
	}

	got := have[:len(have):len(have)]
	gslice.Insert(&got, -2, 1000)

	checkIntSlice(want, got, t)
}
