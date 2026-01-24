package greflect

import (
	"fmt"
	"github.com/skypbc/goutils/gerrors"
	"reflect"
)

// template.FuncMap

func TryEqual(value1 reflect.Value, value2 reflect.Value) (ok bool, err error) {
	value1, value2 = LastElem(value1), LastElem(value2)
	k1, k2 := Kind(value1), Kind(value2)

	if k1 == k2 {
		switch k1 {
		case Bool:
			ok = value1.Bool() == value2.Bool()
		case Int:
			ok = value1.Int() == value2.Int()
		case Uint:
			ok = value1.Uint() == value2.Uint()
		case Float:
			ok = value1.Float() == value2.Float()
		case Complex:
			ok = value1.Complex() == value2.Complex()
		case String:
			ok = value1.String() == value2.String()
		case Nil:
			ok = true
		default:
			if value1.Kind() != value2.Kind() {
				return false, gerrors.NewTypeError().
					SetTemplate("not comparable")
			}
			if !value1.Type().Comparable() {
				return false, gerrors.NewTypeError().
					SetTemplate(`value "{value}" of type "{type}" isn't comparable`).
					AddStr("value", fmt.Sprintf("%s", value1)).
					AddStr("type", fmt.Sprintf("%#+v", value1.Type()))
			}
			ok = value1.Interface() == value2.Interface()
		}
		return ok, nil
	}

	switch {
	case k1 == Int && k2 == Uint:
		ok = value1.Int() >= 0 && uint64(value1.Int()) == value2.Uint()
	case k1 == Uint && k2 == Int:
		ok = value2.Int() >= 0 && uint64(value2.Int()) == value1.Uint()
	default:
		return false, gerrors.NewTypeError().
			SetTemplate(`type cannot be compared: {value1} {value1_type} ~ {value2} {value2_type}`).
			AddStr("value1", fmt.Sprintf("%s", value1)).
			AddStr("value1_type", fmt.Sprintf("%#+v", value1.Type())).
			AddStr("value2", fmt.Sprintf("%s", value2)).
			AddStr("value2_type", fmt.Sprintf("%#+v", value2.Type()))
	}
	return ok, nil
}

func TryLess(value1 reflect.Value, value2 reflect.Value) (ok bool, err error) {
	value1, value2 = LastElem(value1), LastElem(value2)
	k1, k2 := Kind(value1), Kind(value2)

	if k1 == k2 {
		switch k1 {
		case Int:
			ok = value1.Int() < value2.Int()
		case Uint:
			ok = value1.Uint() < value2.Uint()
		case Float:
			ok = value1.Float() < value2.Float()
		case String:
			ok = value1.String() < value2.String()
		default:
			err = gerrors.NewTypeError().
				SetTemplate(`value "{value}" of type "{type}" isn't comparable`).
				AddStr("value", fmt.Sprintf("%s", value1)).
				AddStr("type", fmt.Sprintf("%#+v", value1.Type()))
		}
		return ok, err
	}

	switch {
	case k1 == Int && k2 == Uint:
		ok = value1.Int() < 0 || uint64(value1.Int()) < value2.Uint()
	case k1 == Uint && k2 == Int:
		ok = value2.Int() < 0 || uint64(value2.Int()) < value1.Uint()
	default:
		err = gerrors.NewTypeError().
			SetTemplate(`type cannot be compared: {value1} {value1_type} ~ {value2} {value2_type}`).
			AddStr("value1", fmt.Sprintf("%s", value1)).
			AddStr("value1_type", fmt.Sprintf("%#+v", value1.Type())).
			AddStr("value2", fmt.Sprintf("%s", value2)).
			AddStr("value2_type", fmt.Sprintf("%#+v", value2.Type()))
	}
	return ok, err
}
