package internal

import (
	"github.com/skypbc/goutils/gerrors"
	"github.com/skypbc/goutils/gnum"
	"reflect"
)

func getFromSlice(s reflect.Value, key string) (reflect.Value, error) {
	if k := s.Kind(); k != reflect.Slice && k != reflect.Array {
		return reflect.Value{}, gerrors.NewIncorrectParamsError().
			SetTemplate("Key \"{key}\" isn't slice or array").
			AddStr("key", key)
	}

	index64, err := gnum.ParseInt(key)
	if err != nil {
		return reflect.Value{}, gerrors.NewParseError(err)
	}

	index := int(index64)
	if index < 0 || index >= s.Len() {
		return reflect.Value{}, gerrors.NewNotFoundError().
			SetTemplate("Index \"{index}\" out of range").
			AddInt("index", index)
	}

	res := s.Index(index)
	if !res.IsValid() {
		return reflect.Value{}, gerrors.NewNotFoundError().
			SetTemplate("Key \"{key}\" not found...").
			AddStr("key", key)
	}

	return res, nil
}
