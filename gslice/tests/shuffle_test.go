package tests

import (
	"fmt"
	"github.com/skypbc/goutils/gslice"
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	have := getTestData()
	want := getTestData()
	got := have[:len(have):len(have)]

	gslice.Shuffle(got)

	if len(want) != len(got) {
		t.Error(fmt.Errorf("len(want) != len(got)"))
		return
	}

	if reflect.DeepEqual(want, got) {
		t.Error(fmt.Errorf("must be different, not equal."))
	}
}
