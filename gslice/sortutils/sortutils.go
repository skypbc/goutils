package sortutils

import (
	"github.com/skypbc/goutils/gnum"
	"sort"
)

func Sort[T any](slice []T) {
	switch v := any(slice).(type) {
	case []int:
		sort.Ints(v)
	case []float64:
		sort.Float64s(v)
	case []string:
		sort.Strings(v)
	default:
		SortByFunc(slice, nil, false)
	}
}

func SortByFunc[T any](slice []T, f func(i, j int) bool, reverse bool) {
	sort.Slice(slice, func(i, j int) bool {
		var isLess bool

		if f != nil {
			isLess = f(i, j)
		} else {
			isLess = gnum.Less(slice[i], slice[j])
		}

		if reverse {
			isLess = !isLess
		}

		return isLess
	})
}
