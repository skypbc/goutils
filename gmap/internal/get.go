package internal

import (
	"github.com/skypbc/goutils/gerrors"
	"reflect"
	"strings"

	"golang.org/x/exp/constraints"
)

func Get(data any, key string, sep string) (res any, err error) {
	var ok bool
	for _, key := range strings.Split(key, sep) {
		switch m := data.(type) {
		case map[string]any:
			data, ok = getFromStandardMap(m, key)
		case map[string]int8:
			data, ok = getFromStandardMap(m, key)
		case map[string]int16:
			data, ok = getFromStandardMap(m, key)
		case map[string]int32:
			data, ok = getFromStandardMap(m, key)
		case map[string]int:
			data, ok = getFromStandardMap(m, key)
		case map[string]int64:
			data, ok = getFromStandardMap(m, key)
		case map[string]uint8:
			data, ok = getFromStandardMap(m, key)
		case map[string]uint16:
			data, ok = getFromStandardMap(m, key)
		case map[string]uint32:
			data, ok = getFromStandardMap(m, key)
		case map[string]uint:
			data, ok = getFromStandardMap(m, key)
		case map[string]uint64:
			data, ok = getFromStandardMap(m, key)
		case map[string]float32:
			data, ok = getFromStandardMap(m, key)
		case map[string]float64:
			data, ok = getFromStandardMap(m, key)
		case map[string]string:
			data, ok = getFromStandardMap(m, key)
		case map[string]bool:
			data, ok = getFromStandardMap(m, key)
		default:
			if r, err := getWithReflection(data, key, sep); err == nil {
				data = r.Interface()
				ok = true
			} else {
				ok = false
			}
		}
		if !ok {
			return res, gerrors.NewKeyNotFoundError().
				SetTemplate("Key \"{key}\" not found").
				AddStr("key", key)
		}
	}

	return data, nil
}

func getFromStandardMap[T constraints.Ordered, R any](m map[T]R, key T) (res R, ok bool) {
	val, ok := m[key]
	return val, ok
}

func getWithReflection(data any, key string, sep string) (res reflect.Value, err error) {
	res = reflect.ValueOf(data)
	keys := strings.Split(key, sep)

	for _, key := range keys {
		switch res.Kind() {
		case reflect.Map:
			res, err = getFromMap(res, key)
		case reflect.Slice, reflect.Array:
			res, err = getFromSlice(res, key)
		}
		if err != nil {
			return res, err
		}

		if !res.IsValid() {
			return res, gerrors.NewUnknownError()
		}

		for res.Kind() == reflect.Interface {
			res = res.Elem()
		}
	}

	if !res.IsValid() {
		return res, gerrors.NewUnknownError()
	}

	return res, nil
}
