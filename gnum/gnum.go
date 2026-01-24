package gnum

import (
	"encoding/json"
	"github.com/skypbc/goutils/gerrors"
	"github.com/skypbc/goutils/greflect"
	"math"
	"reflect"
	"strconv"

	cs "golang.org/x/exp/constraints"
)

func ParseInt(text string) (int64, error) {
	res, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return 0, gerrors.NewParseError(err).
			SetTemplate("gnum.ParseInt({text})").
			AddStr("text", text)

	}
	return res, nil
}

func ParseUint(text string) (uint64, error) {
	res, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		return 0, gerrors.NewParseError(err).
			SetTemplate("gnum.ParseUint({text})").
			AddStr("text", text)
	}
	return res, nil
}

func ParseFloat(text string) (float64, error) {
	res, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0, gerrors.NewParseError(err).
			SetTemplate("gnum.ParseFloat({text})").
			AddStr("text", text)
	}
	return res, nil
}

func ParseBool(text string) (bool, error) {
	res, err := strconv.ParseBool(text)
	if err != nil {
		return false, gerrors.NewParseError(err).
			SetTemplate("gnum.ParseBool({text})").
			AddStr("text", text)
	}
	return res, nil
}

func IntTo[T cs.Integer | cs.Float | ~string](value int64) T {
	out, _ := TryIntTo[T](value)
	return out
}

func TryIntTo[T any](value int64) (T, bool) {
	var res T
	switch v := any(&res).(type) {
	case *string:
		*v = strconv.FormatInt(value, 10)
	case *int:
		*v = int(value)
	case *int8:
		*v = int8(value)
	case *int16:
		*v = int16(value)
	case *int32:
		*v = int32(value)
	case *int64:
		*v = int64(value)
	case *uint:
		*v = uint(value)
	case *uint8:
		*v = uint8(value)
	case *uint16:
		*v = uint16(value)
	case *uint32:
		*v = uint32(value)
	case *uint64:
		*v = uint64(value)
	case *float32:
		*v = float32(value)
	case *float64:
		*v = float64(value)
	default:
		x := reflect.ValueOf(value)
		y := reflect.ValueOf(&res).Elem()
		if !x.CanConvert(y.Type()) {
			return res, false
		}
		y.Set(x.Convert(y.Type()))
	}
	return res, true
}

func UintTo[T cs.Integer | cs.Float | ~string](value uint64) T {
	out, _ := TryUintTo[T](value)
	return out
}

func TryUintTo[T any](value uint64) (T, bool) {
	var res T
	switch v := any(&res).(type) {
	case *string:
		*v = strconv.FormatUint(value, 10)
	case *int:
		*v = int(value)
	case *int8:
		*v = int8(value)
	case *int16:
		*v = int16(value)
	case *int32:
		*v = int32(value)
	case *int64:
		*v = int64(value)
	case *uint:
		*v = uint(value)
	case *uint8:
		*v = uint8(value)
	case *uint16:
		*v = uint16(value)
	case *uint32:
		*v = uint32(value)
	case *uint64:
		*v = uint64(value)
	case *float32:
		*v = float32(value)
	case *float64:
		*v = float64(value)
	default:
		x := reflect.ValueOf(value)
		y := reflect.ValueOf(&res).Elem()
		if !x.CanConvert(y.Type()) {
			return res, false
		}
		y.Set(x.Convert(y.Type()))
	}
	return res, true
}

func FloatTo[T cs.Integer | cs.Float | ~string](value float64) T {
	out, _ := TryFloatTo[T](value)
	return out
}

func TryFloatTo[T any](value float64) (T, bool) {
	var res T
	switch v := any(&res).(type) {
	case *string:
		*v = strconv.FormatFloat(value, 'f', -1, 64)
	case *int:
		*v = int(value)
	case *int8:
		*v = int8(value)
	case *int16:
		*v = int16(value)
	case *int32:
		*v = int32(value)
	case *int64:
		*v = int64(value)
	case *uint:
		*v = uint(value)
	case *uint8:
		*v = uint8(value)
	case *uint16:
		*v = uint16(value)
	case *uint32:
		*v = uint32(value)
	case *uint64:
		*v = uint64(value)
	case *float32:
		*v = float32(value)
	case *float64:
		*v = float64(value)
	default:
		x := reflect.ValueOf(value)
		y := reflect.ValueOf(&res).Elem()
		if !x.CanConvert(y.Type()) {
			return res, false
		}
		y.Set(x.Convert(y.Type()))
	}
	return res, true
}

func StringTo[T cs.Integer | cs.Float | ~string](value string) any {
	out, _ := TryStringTo[T](value)
	return out
}

func TryStringTo[T any](value string) (res T, ok bool) {
	if x, ok := TryFloat(value); ok {
		if x != math.Trunc(x) {
			res, ok = TryFloatTo[T](x)
			return res, ok
		}
		return TryIntTo[T](int64(x))
	}
	return res, false
}

func TryAnyTo[T any](value any) (out T, ok bool) {
	switch v := any(value).(type) {
	case string:
		if res, ok := TryFloat(v); ok {
			if res != math.Trunc(res) {
				return TryFloatTo[T](res)
			}
			return TryIntTo[T](int64(res))
		}
	case int, int8, int16, int32, int64:
		if res, ok := TryInt(v); ok {
			return TryIntTo[T](res)
		}
	case uint, uint8, uint16, uint32, uint64, uintptr:
		if res, ok := TryUint(v); ok {
			return TryUintTo[T](res)
		}
	case float32, float64:
		if res, ok := TryFloat(v); ok {
			return TryFloatTo[T](res)
		}
	}

	x := reflect.ValueOf(out)
	y := reflect.ValueOf(value)

	if y.CanConvert(x.Type()) {
		x.Set(y.Convert(x.Type()))
		return out, true
	}

	if bytes, err := json.Marshal(value); err == nil {
		if err = json.Unmarshal(bytes, &out); err == nil {
			return out, true
		}
	}
	return out, false
}

func ToInt[T cs.Integer | cs.Float | ~string](value T) (out int64) {
	out, _ = TryInt(value)
	return out
}

func TryInt(value any) (out int64, ok bool) {
	ok = true
	switch v := any(value).(type) {
	case string:
		var err error
		if out, err = strconv.ParseInt(v, 10, 64); err != nil {
			ok = false
		}
	case int:
		out = int64(v)
	case int8:
		out = int64(v)
	case int16:
		out = int64(v)
	case int32:
		out = int64(v)
	case int64:
		out = v

	case uint:
		out = int64(v)
	case uint8:
		out = int64(v)
	case uint16:
		out = int64(v)
	case uint32:
		out = int64(v)
	case uint64:
		out = int64(v)

	case float64:
		out = int64(v)
	case float32:
		out = int64(v)

	default:
		x := reflect.ValueOf(value)
		switch k := greflect.Kind(x); k {
		case greflect.Int:
			out = x.Int()
		case greflect.Uint:
			out = int64(x.Uint())
		case greflect.Float:
			out = int64(x.Float())
		case greflect.Bool:
			if x.Bool() {
				out = 1
			}
		case greflect.String:
			out, ok = TryStringTo[int64](x.String())
		default:
			ok = false
		}
	}
	return out, ok
}

func TryBool(value any) (out bool, ok bool) {
	ok = true
	switch v := any(value).(type) {
	case string:
		var err error
		if out, err = strconv.ParseBool(v); err != nil {
			ok = false
		}
	case bool:
		out = v
	default:
		var f float64
		if f, ok = TryFloat(value); !ok {
			out = false
			ok = false
		} else if f == 1.0 {
			out = true
		}
	}
	return out, ok
}

func ToUint[T cs.Integer | cs.Float | ~string](value T) (out uint64) {
	out, _ = TryUint(value)
	return out
}

func TryUint(value any) (out uint64, ok bool) {
	ok = true
	switch v := any(value).(type) {
	case string:
		var err error
		if out, err = strconv.ParseUint(v, 10, 64); err != nil {
			ok = false
		}
	case uint:
		out = uint64(v)
	case uint8:
		out = uint64(v)
	case uint16:
		out = uint64(v)
	case uint32:
		out = uint64(v)
	case uint64:
		out = v

	case int:
		out = uint64(v)
	case int8:
		out = uint64(v)
	case int16:
		out = uint64(v)
	case int32:
		out = uint64(v)
	case int64:
		out = uint64(v)

	case float64:
		out = uint64(v)
	case float32:
		out = uint64(v)

	default:
		x := reflect.ValueOf(value)
		switch k := greflect.Kind(x); k {
		case greflect.Int:
			out = uint64(x.Int())
		case greflect.Uint:
			out = x.Uint()
		case greflect.Float:
			out = uint64(x.Float())
		case greflect.Bool:
			if x.Bool() {
				out = 1
			}
		case greflect.String:
			out, ok = TryStringTo[uint64](x.String())
		default:
			ok = false
		}
	}
	return out, ok
}

func ToFloat[T cs.Integer | cs.Float | ~string](value T) (out float64) {
	out, _ = TryFloat(value)
	return out
}

func TryFloat(value any) (out float64, ok bool) {
	ok = true
	switch v := any(value).(type) {
	case string:
		var err error
		if out, err = strconv.ParseFloat(v, 64); err != nil {
			ok = false
		}
	case float64:
		out = v
	case float32:
		out = float64(v)

	case int:
		out = float64(v)
	case int8:
		out = float64(v)
	case int16:
		out = float64(v)
	case int32:
		out = float64(v)
	case int64:
		out = float64(v)

	case uint:
		out = float64(v)
	case uint8:
		out = float64(v)
	case uint16:
		out = float64(v)
	case uint32:
		out = float64(v)
	case uint64:
		out = float64(v)

	default:
		x := reflect.ValueOf(value)
		switch k := greflect.Kind(x); k {
		case greflect.Int:
			out = float64(x.Int())
		case greflect.Uint:
			out = float64(x.Uint())
		case greflect.Float:
			out = x.Float()
		case greflect.Bool:
			if x.Bool() {
				out = 1.0
			}
		case greflect.String:
			out, ok = TryStringTo[float64](x.String())
		default:
			ok = false
		}
	}
	return out, ok
}

func ToString(value any) string {
	out, _ := TryAnyTo[string](value)
	return out
}

func TryString(value any) (out string, ok bool) {
	return TryAnyTo[string](value)
}

func TryOut(value any, out any) (ok bool) {
	switch v := out.(type) {
	case *string:
		if x, ok := TryAnyTo[string](value); ok {
			*v = x
			return true
		}
	case *int:
		if x, ok := TryInt(value); ok {
			*v = int(x)
			return true
		}
	case *int8:
		if x, ok := TryInt(value); ok {
			*v = int8(x)
			return true
		}
	case *int16:
		if x, ok := TryInt(value); ok {
			*v = int16(x)
			return true
		}
	case *int32:
		if x, ok := TryInt(value); ok {
			*v = int32(x)
			return true
		}
	case *int64:
		if x, ok := TryInt(value); ok {
			*v = x
			return true
		}
	case *uint:
		if x, ok := TryUint(value); ok {
			*v = uint(x)
			return true
		}
	case *uint8:
		if x, ok := TryUint(value); ok {
			*v = uint8(x)
			return true
		}
	case *uint16:
		if x, ok := TryUint(value); ok {
			*v = uint16(x)
			return true
		}
	case *uint32:
		if x, ok := TryUint(value); ok {
			*v = uint32(x)
			return true
		}
	case *uint64:
		if x, ok := TryUint(value); ok {
			*v = uint64(x)
			return true
		}
	case *uintptr:
		if x, ok := TryUint(value); ok {
			*v = uintptr(x)
			return true
		}
	case *float32:
		if x, ok := TryFloat(value); ok {
			*v = float32(x)
			return true
		}
	case *float64:
		if x, ok := TryFloat(value); ok {
			*v = x
			return true
		}
	}

	x := reflect.ValueOf(value)
	y := reflect.ValueOf(out)

	if y.Kind() != reflect.Pointer {
		return false
	}

	y = y.Elem()
	yt := y.Type()

	if x.CanConvert(yt) {
		y.Set(x.Convert(yt))
		return true
	}

	if bytes, err := json.Marshal(value); err == nil {
		if err = json.Unmarshal(bytes, out); err == nil {
			return true
		}
	}
	return false
}
