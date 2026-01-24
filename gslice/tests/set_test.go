package tests

import (
	"fmt"
	"github.com/skypbc/goutils/gslice"
	"testing"
)

// go test -run TestSet -v

func TestSetAtStart1(t *testing.T) {
	have := getTestData()
	want := []int{
		1000, 2, 3, 4, 5,
	}

	got := have[:len(have):len(have)]
	gslice.Set(got, 0, 1000)

	checkIntSlice(want, got, t)
}

func TestSetAtStart2(t *testing.T) {
	have := getTestData()
	want := []int{
		1000, 2, 3, 4, 5,
	}

	got := have[:len(have):len(have)]
	gslice.Set(got, -5, 1000)

	checkIntSlice(want, got, t)
}

func TestSetAtEnd1(t *testing.T) {
	have := getTestData()
	want := []int{
		1, 2, 3, 4, 1000,
	}

	got := have[:len(have):len(have)]
	gslice.Set(got, len(got)-1, 1000)

	checkIntSlice(want, got, t)
}

func TestSetAtEnd2(t *testing.T) {
	have := getTestData()
	want := []int{
		1, 2, 3, 4, 1000,
	}

	got := have[:len(have):len(have)]
	gslice.Set(got, -1, 1000)

	checkIntSlice(want, got, t)
}

func TestSetWithOutOfRange1(t *testing.T) {
	have := getTestData()
	got := have[:len(have):len(have)]

	err := gslice.TrySet(got, len(got), 1000)

	if err == nil {
		t.Error(fmt.Errorf("want err != nil, got err == nil..."))
	}
}

func TestSetWithOutOfRange2(t *testing.T) {
	have := getTestData()
	got := have[:len(have):len(have)]

	err := gslice.TrySet(got, 1000, 1000)

	if err == nil {
		t.Error(fmt.Errorf("want err != nil, got err == nil..."))
	}
}

func TestSetWithOutOfRange3(t *testing.T) {
	have := getTestData()
	got := have[:len(have):len(have)]

	err := gslice.TrySet(got, -6, 1000)

	if err == nil {
		t.Error(fmt.Errorf("want err != nil, got err == nil..."))
	}
}

func TestSetWithOutOfRange4(t *testing.T) {
	have := getTestData()
	got := have[:len(have):len(have)]

	err := gslice.TrySet(got, -1000, 1000)

	if err == nil {
		t.Error(fmt.Errorf("want err != nil, got err == nil..."))
	}
}
