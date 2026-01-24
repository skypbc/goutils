package gbytes

import (
	"github.com/skypbc/goutils/gerrors"

	"golang.org/x/exp/constraints"
)

func Set[T constraints.Integer](data []byte, offset int, value T, width int) int {
	offset, err := TrySet(data, offset, value, width)
	if err != nil {
		panic(err)
	}
	return offset
}

func TrySet[T constraints.Integer](data []byte, offset int, value T, width int) (int, error) {
	if width > 8 {
		return 0, gerrors.NewIncorrectParamsError().
			SetTemplate(`invalid width: "{width}", must be <= 8`).
			AddInt("width", width)
	}

	switch v := uint64(value); width {
	case 0:
	case 1:
		data[offset] = uint8(v)
	case 2:
		data[offset+0] = uint8((v & 0x00_FF))
		data[offset+1] = uint8((v & 0xFF_00) >> 8)
	case 3:
		data[offset+0] = uint8((v & 0x00_00_00_FF))
		data[offset+1] = uint8((v & 0x00_00_FF_00) >> 8)
		data[offset+2] = uint8((v & 0x00_FF_00_00) >> 16)
	case 4:
		data[offset+0] = uint8((v & 0x00_00_00_FF))
		data[offset+1] = uint8((v & 0x00_00_FF_00) >> 8)
		data[offset+2] = uint8((v & 0x00_FF_00_00) >> 16)
		data[offset+3] = uint8((v & 0xFF_00_00_00) >> 24)
	case 5:
		data[offset+0] = uint8((v & 0x00_00_00_00_00_00_00_FF))
		data[offset+1] = uint8((v & 0x00_00_00_00_00_00_FF_00) >> 8)
		data[offset+2] = uint8((v & 0x00_00_00_00_00_FF_00_00) >> 16)
		data[offset+3] = uint8((v & 0x00_00_00_00_FF_00_00_00) >> 24)
		data[offset+4] = uint8((v & 0x00_00_00_FF_00_00_00_00) >> 32)
	case 6:
		data[offset+0] = uint8((v & 0x00_00_00_00_00_00_00_FF))
		data[offset+1] = uint8((v & 0x00_00_00_00_00_00_FF_00) >> 8)
		data[offset+2] = uint8((v & 0x00_00_00_00_00_FF_00_00) >> 16)
		data[offset+3] = uint8((v & 0x00_00_00_00_FF_00_00_00) >> 24)
		data[offset+4] = uint8((v & 0x00_00_00_FF_00_00_00_00) >> 32)
		data[offset+5] = uint8((v & 0x00_00_FF_00_00_00_00_00) >> 40)
	case 7:
		data[offset+0] = uint8((v & 0x00_00_00_00_00_00_00_FF))
		data[offset+1] = uint8((v & 0x00_00_00_00_00_00_FF_00) >> 8)
		data[offset+2] = uint8((v & 0x00_00_00_00_00_FF_00_00) >> 16)
		data[offset+3] = uint8((v & 0x00_00_00_00_FF_00_00_00) >> 24)
		data[offset+4] = uint8((v & 0x00_00_00_FF_00_00_00_00) >> 32)
		data[offset+5] = uint8((v & 0x00_00_FF_00_00_00_00_00) >> 40)
		data[offset+6] = uint8((v & 0x00_FF_00_00_00_00_00_00) >> 48)
	default:
		data[offset+0] = uint8((v & 0x00_00_00_00_00_00_00_FF))
		data[offset+1] = uint8((v & 0x00_00_00_00_00_00_FF_00) >> 8)
		data[offset+2] = uint8((v & 0x00_00_00_00_00_FF_00_00) >> 16)
		data[offset+3] = uint8((v & 0x00_00_00_00_FF_00_00_00) >> 24)
		data[offset+4] = uint8((v & 0x00_00_00_FF_00_00_00_00) >> 32)
		data[offset+5] = uint8((v & 0x00_00_FF_00_00_00_00_00) >> 40)
		data[offset+6] = uint8((v & 0x00_FF_00_00_00_00_00_00) >> 48)
		data[offset+7] = uint8((v & 0xFF_00_00_00_00_00_00_00) >> 56)
	}

	return offset + width, nil
}

func SetByte[T constraints.Integer](data []byte, offset int, value T) int {
	return Set(data, offset, value, 1)
}

func SetWord[T constraints.Integer](data []byte, offset int, value T) int {
	return Set(data, offset, value, 2)
}

func SetDword[T constraints.Integer](data []byte, offset int, value T) int {
	return Set(data, offset, value, 4)
}

func SetQword[T constraints.Integer](data []byte, offset int, value T) int {
	return Set(data, offset, value, 8)
}
