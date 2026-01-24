package gnum

import (
	"github.com/skypbc/goutils/gerrors"
	"reflect"
)

func Equal[T2 any, T1 any](a T1, b T2) bool {
	res, err := TryEqual(a, b)
	if err != nil {
		panic(err)
	}
	return res
}

func TryEqual[T2 any, T1 any](a T1, b T2) (bool, error) {
	x := reflect.ValueOf(a)
	if x.Comparable() {
		y := reflect.ValueOf(b)
		xt, yt := x.Type(), y.Type()
		if xt != yt {
			if x.CanConvert(yt) {
				x = x.Convert(yt)
				return x.Equal(y), nil
			}
			return false, gerrors.NewTypeError().
				SetTemplate("type cannot be compared")
		}
		return x.Equal(y), nil
	}
	return false, gerrors.NewTypeError().
		SetTemplate("type cannot be compared")
}
