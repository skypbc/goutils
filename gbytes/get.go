package gbytes

import (
	"github.com/skypbc/goutils/gerrors"

	"golang.org/x/exp/constraints"
)

func Get[T constraints.Integer](data []byte, offset int, width int) (res T, off int) {
	if width > 8 {
		err := gerrors.NewIncorrectParamsError().
			SetTemplate(`invalid width: "{width}", must be <= 8`).
			AddInt("width", width)
		panic(err)
	}
	var val uint64
	for i, bt := range data[offset : offset+width] {
		val |= uint64(bt) << (i * 8)
	}
	return T(val), offset + width
}

func Get2[T constraints.Integer](data []byte, offset int, width int) T {
	res, _ := Get[T](data, offset, width)
	return res
}

func TryGet[T constraints.Integer](data []byte, offset int, width int) (res T, off int, err error) {
	if ((len(data)) < offset+width) || width > 8 {
		return res, offset, gerrors.NewIncorrectParamsError().
			SetTemplate(`invalid parameters: data length="{data_length}", offset="{offset}", width="{width}"`).
			AddInt("data_length", len(data)).
			AddInt("offset", offset).
			AddInt("width", width)
	}
	var val uint64
	for i, bt := range data[offset : offset+width] {
		val |= uint64(bt) << (i * 8)
	}
	return T(val), offset + width, nil
}

func TryGet2[T constraints.Integer](data []byte, offset int, width int) (res T, err error) {
	res, _, err = TryGet[T](data, offset, width)
	return res, err
}
