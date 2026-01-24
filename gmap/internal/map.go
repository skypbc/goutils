package internal

import (
	"fmt"
	"github.com/skypbc/goutils/gerrors"
	"github.com/skypbc/goutils/gnum"
	"github.com/skypbc/goutils/greflect"
	"reflect"
)

func getFromMap(m reflect.Value, key string) (value reflect.Value, err error) {
	if m.Kind() != reflect.Map {
		return reflect.Value{}, gerrors.NewIncorrectParamsError().
			SetTemplate("Key \"{key}\" isn't map").
			AddStr("key", key)
	}

	mapKey, err := convertKey(m.Type().Key(), key)
	if err != nil {
		return mapKeyBrutforce(m, key)
	}

	if value = m.MapIndex(mapKey); !value.IsValid() {
		return reflect.Value{}, gerrors.NewNotFoundError().
			SetTemplate("Key \"{key}\" not found...").
			AddStr("key", key)
	}

	return value, nil
}

func convertKey(keyType reflect.Type, key string) (reflect.Value, error) {
	switch k := greflect.Kind2(keyType); k {
	case greflect.String:
		rKey := reflect.ValueOf(key)
		if !rKey.CanConvert(keyType) {
			return reflect.Value{}, gerrors.NewIncorrectParamsError().
				SetTemplate(`cannot convert key "{key}" to type "{type}"`).
				AddStr("key", key).
				AddStr("type", keyType.String())

		}
		return rKey.Convert(keyType), nil

	case greflect.Int, greflect.Uint:
		i64, err := gnum.ParseInt(key)
		if err != nil {
			return reflect.Value{}, err
		}
		rKey := reflect.ValueOf(i64)
		if !rKey.CanConvert(keyType) {
			return reflect.Value{}, gerrors.NewIncorrectParamsError().
				SetTemplate("cannot convert key \"{key}\" to type \"{type}\"").
				AddStr("key", key).
				AddStr("type", keyType.String())
		}
		return rKey.Convert(keyType), nil

	default:
		return reflect.Value{}, gerrors.NewIncorrectParamsError().
			SetTemplate(`unsupported key "{key}" type "{type}"`).
			AddStr("key", key).
			AddStr("type", keyType.String())
	}
}

func mapKeyBrutforce(m reflect.Value, key string) (reflect.Value, error) {
	for _, res := range m.MapKeys() {
		if fmt.Sprintf("%v", res.Interface()) == key {
			return m.MapIndex(res), nil
		}
	}
	return reflect.Value{}, gerrors.NewParseError()
}
