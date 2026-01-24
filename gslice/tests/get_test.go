package tests

import (
	"fmt"
	"github.com/skypbc/goutils/gslice"
	"testing"
)

// go test -run TestGet -v

func TestGetFromStart1(t *testing.T) {
	have := getTestData()
	want := 1
	got := gslice.Get(have, 0)

	if want != got {
		t.Error(fmt.Errorf("want != got, (%d != %d)", want, got))
	}
}

func TestGetFromStart2(t *testing.T) {
	have := getTestData()
	want := 1
	got := gslice.Get(have, -5)

	if want != got {
		t.Error(fmt.Errorf("want != got, (%d != %d)", want, got))
	}
}

func TestGetFromEnd1(t *testing.T) {
	have := getTestData()
	want := 5
	got := gslice.Get(have, len(have)-1)

	if want != got {
		t.Error(fmt.Errorf("want != got, (%d != %d)", want, got))
	}
}

func TestGetFromEnd2(t *testing.T) {
	have := getTestData()
	want := 5
	got := gslice.Get(have, -1)

	if want != got {
		t.Error(fmt.Errorf("want != got, (%d != %d)", want, got))
	}
}

func TestGetWithOutOfRange1(t *testing.T) {
	have := getTestData()

	_, err := gslice.TryGet(have, len(have))
	if err == nil {
		t.Error(fmt.Errorf("want err != nil, got err == nil..."))
	}
}

func TestGetWithOutOfRange2(t *testing.T) {
	have := getTestData()

	_, err := gslice.TryGet(have, 1000)
	if err == nil {
		t.Error(fmt.Errorf("want err != nil, got err == nil..."))
	}
}

func TestGetWithOutOfRange3(t *testing.T) {
	have := getTestData()

	_, err := gslice.TryGet(have, -6)
	if err == nil {
		t.Error(fmt.Errorf("want err != nil, got err == nil..."))
	}
}

func TestGetWithOutOfRange4(t *testing.T) {
	have := getTestData()

	_, err := gslice.TryGet(have, -1000)
	if err == nil {
		t.Error(fmt.Errorf("want err != nil, got err == nil..."))
	}
}
