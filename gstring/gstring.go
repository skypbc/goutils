package gstring

import (
	"fmt"
	"github.com/skypbc/goutils/gfmt"
	"github.com/skypbc/goutils/gnum"
	"github.com/skypbc/goutils/greflect"
	"reflect"
	"strings"
)

func StartsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func EndsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

func Reverse(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func ReplaceFromEnd(path string, old string, new string, n int) string {
	return Reverse(strings.Replace(Reverse(path), Reverse(old), Reverse(new), n))
}

func String(val any) string {
	if s, ok := TryString(val); ok {
		return s
	}
	return ""
}

func TryString(val any) (res string, ok bool) {
	switch v := val.(type) {
	case string:
		return v, true
	case fmt.Stringer:
		return v.String(), true
	}

	// Рефлексия используется для обработки производных типов от типа "string".
	rResType := reflect.TypeOf(res)
	rValue := reflect.ValueOf(val)

	switch greflect.Kind(rValue) {
	case greflect.String:
		if rValue.CanConvert(rResType) {
			return rValue.Convert(rResType).Interface().(string), true
		}
	case greflect.Int, greflect.Uint, greflect.Float:
		return gnum.TryAnyTo[string](val)
	case greflect.Bool:
		return gfmt.Sprintf("%v", val), true
	}

	return "", false
}
