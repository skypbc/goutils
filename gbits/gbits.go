package gbits

import (
	cs "golang.org/x/exp/constraints"
)

func Set[T cs.Integer](val T, pos int) T {
	return val | (1 << T(pos))
}

func Clear[T cs.Integer](val T, pos int) T {
	return val & T(^(1 << pos))
}

func Is[T cs.Integer](val T, pos int) bool {
	return (val & (1 << pos)) != 0
}

func Pop[T cs.Integer](val T, pos int) (T, bool) {
	return Clear(val, pos), Is(val, pos)
}

func Range[T cs.Integer](val T, from int, to int) T {
	return (val >> from) & (1<<(to-from) - 1)
}

func Range2[T cs.Integer](val T, from int, width int) T {
	if width == 0 {
		return val
	}
	return (val >> from) & (1<<width - 1)
}

func SetRange[T cs.Integer](val T, from int, to int) T {
	mask := T((1<<(to-from))-1) << from
	return val | mask
}

func SetRange2[T cs.Integer](val T, from int, width int) T {
	if width == 0 {
		return val
	}
	mask := T((1<<width)-1) << from
	return val | mask
}

func SetRangeFrom[T cs.Integer](val T, pos int, rng T) T {
	return val | (rng << pos)
}

func ClearRange[T cs.Integer](val T, from int, to int) T {
	mask := T((1<<(to-from))-1) << from
	return val & ^mask
}

func ClearRange2[T cs.Integer](val T, from int, width int) T {
	if width == 0 {
		return val
	}
	mask := T((1<<width)-1) << from
	return val & ^mask
}
