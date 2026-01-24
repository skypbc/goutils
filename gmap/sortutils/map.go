package sortutils

import (
	"github.com/skypbc/goutils/gnum"
	"github.com/skypbc/goutils/gtypes"
	"sort"
)

func SortByFunc[M ~map[K]V, K comparable, V any](
	m M,
	f func(p1, p2 gtypes.MapItem[K, V]) bool,
	reverse bool,
	byVal bool,
) []gtypes.MapItem[K, V] {
	res := make([]gtypes.MapItem[K, V], len(m))

	index := 0
	for key, val := range m {
		res[index] = gtypes.MapItem[K, V]{Key: key, Value: val}
		index++
	}

	sort.Slice(res, func(i, j int) bool {
		var isLess bool

		if f != nil {
			isLess = f(res[i], res[j])
		} else if !byVal {
			isLess = gnum.Less(any(res[i].Key), any(res[j].Key))
		} else {
			isLess = gnum.Less(any(res[i].Value), any(res[j].Value))
		}

		if reverse {
			isLess = !isLess
		}

		return isLess
	})

	return res
}
