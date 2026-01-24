package gbytes

import "golang.org/x/exp/constraints"

func Width[T constraints.Integer](val T) int {
	switch v := uint64(val); {
	case v <= 0xFF:
		return 1
	case v <= 0xFF_FF:
		return 2
	case v <= 0xFF_FF_FF:
		return 3
	case v <= 0xFF_FF_FF_FF:
		return 4
	case v <= 0xFF_FF_FF_FF_FF:
		return 5
	case v <= 0xFF_FF_FF_FF_FF_FF:
		return 6
	case v <= 0xFF_FF_FF_FF_FF_FF_FF:
		return 7
	default:
		return 8
	}
}

func Width2[T constraints.Integer | constraints.Float](val T) int {
	switch any(val).(type) {
	case int8, uint8:
		return 1
	case int16, uint16:
		return 2
	case int32, uint32:
		return 4
	case int64, uint64:
		return 8
	case float32:
		return 4
	case float64:
		return 8
	default:
		return 0
	}
}
