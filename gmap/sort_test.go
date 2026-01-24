package gmap_test

import (
	"fmt"
	"github.com/skypbc/goutils/gmap"
	"github.com/skypbc/goutils/gtypes"
	"testing"
)

var mapTestData map[string]int = map[string]int{
	"a": 3, "b": 1, "c": 5,
	"d": 0, "e": 2, "f": 4,
}

func TestSortMapByKeyDefault(t *testing.T) {
	want := []gtypes.MapItem[string, int]{
		{Key: "a", Value: 3}, {Key: "b", Value: 1}, {Key: "c", Value: 5},
		{Key: "d", Value: 0}, {Key: "e", Value: 2}, {Key: "f", Value: 4},
	}
	got := gmap.Sort(mapTestData)

	checkMap(want, got, t)
}

func TestSortMapByKey(t *testing.T) {
	want := []gtypes.MapItem[string, int]{
		{Key: "a", Value: 3}, {Key: "b", Value: 1}, {Key: "c", Value: 5},
		{Key: "d", Value: 0}, {Key: "e", Value: 2}, {Key: "f", Value: 4},
	}
	got := gmap.SortByKey(mapTestData)

	checkMap(want, got, t)
}

func TestSortMapByKeyReverse(t *testing.T) {
	want := []gtypes.MapItem[string, int]{
		{Key: "f", Value: 4}, {Key: "e", Value: 2}, {Key: "d", Value: 0},
		{Key: "c", Value: 5}, {Key: "b", Value: 1}, {Key: "a", Value: 3},
	}
	got := gmap.SortByKeyReverse(mapTestData)

	checkMap(want, got, t)
}

func TestSortMapByValue(t *testing.T) {
	want := []gtypes.MapItem[string, int]{
		{Key: "d", Value: 0}, {Key: "b", Value: 1}, {Key: "e", Value: 2},
		{Key: "a", Value: 3}, {Key: "f", Value: 4}, {Key: "c", Value: 5},
	}
	got := gmap.SortByValue(mapTestData)

	checkMap(want, got, t)
}

func TestSortMapByValueReverse(t *testing.T) {
	want := []gtypes.MapItem[string, int]{
		{Key: "c", Value: 5}, {Key: "f", Value: 4}, {Key: "a", Value: 3},
		{Key: "e", Value: 2}, {Key: "b", Value: 1}, {Key: "d", Value: 0},
	}
	got := gmap.SortByValueReverse(mapTestData)

	checkMap(want, got, t)
}

func TestSortMapWith(t *testing.T) {
	want := []gtypes.MapItem[string, int]{
		{Key: "d", Value: 0}, {Key: "b", Value: 1}, {Key: "e", Value: 2},
		{Key: "a", Value: 3}, {Key: "f", Value: 4}, {Key: "c", Value: 5},
	}

	got := gmap.SortByFunc(mapTestData, func(p1, p2 gtypes.MapItem[string, int]) bool {
		return p1.Value < p2.Value
	})

	checkMap(want, got, t)
}

func TestSortMapWithReverse(t *testing.T) {
	want := []gtypes.MapItem[string, int]{
		{Key: "c", Value: 5}, {Key: "f", Value: 4}, {Key: "a", Value: 3},
		{Key: "e", Value: 2}, {Key: "b", Value: 1}, {Key: "d", Value: 0},
	}

	got := gmap.SortByFuncReverse(mapTestData, func(p1, p2 gtypes.MapItem[string, int]) bool {
		return p1.Value < p2.Value
	})

	checkMap(want, got, t)
}

func checkMap(want, got gtypes.MapItemList[string, int], t *testing.T) {
	err := checkMapEx(want, got)
	if err != nil {
		t.Error(err)
	}
}

func checkMapEx(want, got gtypes.MapItemList[string, int]) error {
	if len(want) != len(got) {
		return fmt.Errorf("len(want) != len(got)")
	}

	for i := 0; i < len(want); i++ {
		w := want[i]
		g := got[i]

		if w.Key != g.Key {
			return fmt.Errorf("Index: %v, Key (%v != %v)", i, w.Key, g.Key)
		}

		if w.Value != g.Value {
			return fmt.Errorf("Index: %v, Key: %v, Value (%v != %v)", i, w.Key, w.Value, g.Value)
		}
	}

	return nil
}
