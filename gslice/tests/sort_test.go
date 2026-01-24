package tests

import (
	"fmt"
	"github.com/skypbc/goutils/gslice"
	"testing"
)

type item struct {
	value int
}

type itemList []item

var sortItems []item = itemList{
	{5}, {1}, {3}, {4}, {2},
}

func TestSortSliceDefault(t *testing.T) {
	want := []int{
		1, 2, 3, 4, 5,
	}

	got := []int{
		5, 1, 3, 4, 2,
	}

	gslice.Sort(got)

	checkIntSlice(want, got, t)
}

func TestSortSliceDefaultReverse(t *testing.T) {
	want := []int{
		5, 4, 3, 2, 1,
	}

	got := []int{
		5, 1, 3, 4, 2,
	}

	gslice.SortReverse(got)

	checkIntSlice(want, got, t)
}

func TestSortSliceByFunc(t *testing.T) {
	want := []item{
		{1}, {2}, {3}, {4}, {5},
	}

	got := append([]item{}, sortItems...)

	gslice.SortByFunc(got, func(i, j int) bool {
		return got[i].value < got[j].value
	})

	checkSlice(want, got, t)
}

func TestSortSliceByFuncReverse(t *testing.T) {
	want := []item{
		{5}, {4}, {3}, {2}, {1},
	}

	got := append([]item{}, sortItems...)

	gslice.SortByFuncReverse(got, func(i, j int) bool {
		return got[i].value < got[j].value
	})

	checkSlice(want, got, t)
}

func checkIntSlice(want []int, got []int, t *testing.T) {
	err := checkIntSliceEx(want, got)
	if err != nil {
		t.Error(err)
	}
}

func checkIntSliceEx(want, got []int) error {
	if len(want) != len(got) {
		return fmt.Errorf("len(want) != len(got)")
	}

	for i := 0; i < len(want); i++ {
		w, g := want[i], got[i]
		if w != g {
			return fmt.Errorf("Index: %v, (%v != %v)", i, w, g)
		}
	}

	return nil
}

func checkSlice(want, got itemList, t *testing.T) {
	err := checkSliceEx(want, got)
	if err != nil {
		t.Error(err)
	}
}

func checkSliceEx(want, got itemList) error {
	if len(want) != len(got) {
		return fmt.Errorf("len(want) != len(got)")
	}

	for i := 0; i < len(want); i++ {
		w := want[i]
		g := got[i]

		if w.value != g.value {
			return fmt.Errorf("Index: %v, (%v != %v)", i, w, g)
		}
	}

	return nil
}
