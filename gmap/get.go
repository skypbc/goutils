package gmap

import (
	"fmt"
	"github.com/skypbc/goutils/gerrors"
	"github.com/skypbc/goutils/gmap/internal"
	"github.com/skypbc/goutils/gnum"
	"reflect"
)

func Get[R any](m any, key string) (R, error) {
	return GetWithSep[R](m, key, ".")
}

func MustGet[R any](m any, key string) R {
	res, err := GetWithSep[R](m, key, ".")
	if err != nil {
		panic(err)
	}
	return res
}

func GetAny(m any, key string) (any, bool) {
	return GetAnyWithSep(m, key, ".")
}

func GetAnyWithSep(m any, key string, sep string) (res any, ok bool) {
	res, err := internal.Get(m, key, sep)
	if err != nil {
		return res, false
	}
	return res, true
}

func GetWithSep[R any](m any, key string, sep string) (res R, err error) {
	raw, err := internal.Get(m, key, sep)
	if err != nil {
		return res, err
	}

	if res, ok := raw.(R); ok {
		return res, nil
	}
	if res, ok := gnum.TryAnyTo[R](raw); ok {
		return res, nil
	}

	rRes := reflect.ValueOf(res)
	rRaw := reflect.ValueOf(raw)
	if rRaw.CanConvert(rRes.Type()) {
		return rRaw.Convert(rRes.Type()).Interface().(R), nil
	}

	return res, gerrors.NewParseError().
		SetTemplate(`Can't convert the "{key}" key from "{value}" value to "{type}" type`).
		AddStr("key", key).
		AddAny("value", fmt.Sprintf("%#v", raw)).
		AddStr("type", fmt.Sprintf("%T", res))
}

func GetOrDefault[R any](m any, key string, default_ R) R {
	result, err := GetWithSep[R](m, key, ".")
	if err != nil {
		return default_
	}
	return result
}

func GetAnyOrDefault(m any, key string, default_ any) any {
	result, err := GetWithSep[any](m, key, ".")
	if err != nil {
		return default_
	}
	return result
}

func GetOrDefaultWithSep[R any](m any, key string, sep string, default_ R) R {
	result, err := GetWithSep[R](m, key, sep)
	if err != nil {
		return default_
	}
	return result
}

func GetAnyOrDefaultWithSep(m any, key string, sep string, default_ any) any {
	result, err := GetWithSep[any](m, key, sep)
	if err != nil {
		return default_
	}
	return result
}
