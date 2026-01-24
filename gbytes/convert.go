package gbytes

import (
	"github.com/skypbc/goutils/gerrors"
	"github.com/skypbc/goutils/ghex"
	"math"

	"golang.org/x/exp/constraints"
)

func FromHex(val string, sep ...string) []byte {
	res, err := TryFromHex(val)
	if err != nil {
		panic(err)
	}
	return res
}

func TryFromHex(val string, sep ...string) ([]byte, error) {
	return ghex.TryToBytes(val, sep...)
}

func From[T constraints.Integer](val T, width ...int) []byte {
	res, err := TryFrom(val, width...)
	if err != nil {
		panic(err)
	}
	return res
}

func TryFrom[T constraints.Integer](val T, width ...int) (res []byte, err error) {
	w := 0
	if len(width) > 0 {
		w = width[0]
	} else {
		w = Width2(val)
	}
	if w == 0 {
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate(`invalid width: "{width}"`).
			AddInt("width", w)
	}
	res = make([]byte, w)
	if _, err = TrySet(res, 0, val, w); err != nil {
		return nil, err
	}
	return res, nil
}

func FromFloat[T constraints.Float](val T, width ...int) []byte {
	res, err := TryFromFloat(val, width...)
	if err != nil {
		panic(err)
	}
	return res
}

func TryFromFloat[T constraints.Float](val T, width ...int) (res []byte, err error) {
	w := 0
	if len(width) > 0 {
		w = width[0]
	} else {
		w = Width2(val)
	}
	switch w {
	case 4:
		return TryFrom(math.Float32bits(float32(val)), w)
	case 8:
		return TryFrom(math.Float64bits(float64(val)), w)
	default:
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate(`invalid width: "{width}"`).
			AddInt("width", w)
	}
}
