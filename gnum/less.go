package gnum

import (
	"fmt"
	"github.com/skypbc/goutils/gerrors"
	"github.com/skypbc/goutils/greflect"
	"reflect"
)

func Less(v1 any, v2 any) bool {
	res, err := TryLess(v1, v2)
	if err != nil {
		panic(err)
	}
	return res
}

func TryLess(v1 any, v2 any) (bool, error) {
	switch v1_ := v1.(type) {

	case string:
		v2_, ok := v2.(string)
		if ok {
			return v1_ < v2_, nil
		}

	case int:
		v2_, ok := v2.(int)
		if ok {
			return v1_ < v2_, nil
		}
	case int8:
		v2_, ok := v2.(int8)
		if ok {
			return v1_ < v2_, nil
		}
	case int16:
		v2_, ok := v2.(int16)
		if ok {
			return v1_ < v2_, nil
		}
	case int32:
		v2_, ok := v2.(int32)
		if ok {
			return v1_ < v2_, nil
		}
	case int64:
		v2_, ok := v2.(int64)
		if ok {
			return v1_ < v2_, nil
		}

	case uint:
		v2_, ok := v2.(uint)
		if ok {
			return v1_ < v2_, nil
		}
	case uint8:
		v2_, ok := v2.(uint8)
		if ok {
			return v1_ < v2_, nil
		}
	case uint16:
		v2_, ok := v2.(uint16)
		if ok {
			return v1_ < v2_, nil
		}
	case uint32:
		v2_, ok := v2.(uint32)
		if ok {
			return v1_ < v2_, nil
		}
	case uint64:
		v2_, ok := v2.(uint64)
		if ok {
			return v1_ < v2_, nil
		}

	default:
		ok, err := greflect.TryLess(reflect.ValueOf(v1), reflect.ValueOf(v2))
		if err != nil {
			return false, gerrors.NewTypeError(err).
				SetTemplate("Unsupported \"{type}\" type...").
				AddStr("type", fmt.Sprintf("%#v", v1))
		}
		return ok, nil
	}

	return false, gerrors.NewTypeError().
		SetTemplate("Inconsistent types: \"{v1}\" - \"{v2}\"").
		AddStr("v1", fmt.Sprintf("%#v", v1)).
		AddStr("v2", fmt.Sprintf("%#v", v2))
}
