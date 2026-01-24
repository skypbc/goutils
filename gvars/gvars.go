package gvars

import (
	"encoding/json"
	"github.com/skypbc/goutils/gerrors"
)

func Unused(...any) {
}

func PointerTo[T any](v T) *T {
	return &v
}

func ConvertTo[T any](from any) (res T, err error) {
	data, ok := from.([]byte)
	if !ok {
		if data, err = json.Marshal(from); err != nil {
			return res, gerrors.Wrap(err)
		}
	}
	if err = json.Unmarshal(data, &res); err != nil {
		return res, gerrors.Wrap(err)
	}
	return res, nil
}

func ConvertTo2[T any](from any, to *T) (err error) {
	data, ok := from.([]byte)
	if !ok {
		if data, err = json.Marshal(from); err != nil {
			return gerrors.Wrap(err)
		}
	}
	return json.Unmarshal(data, to)
}

func ConvertTo3[T any](from any, to T) (err error) {
	data, ok := from.([]byte)
	if !ok {
		if data, err = json.Marshal(from); err != nil {
			return gerrors.Wrap(err)
		}
	}
	return json.Unmarshal(data, to)
}
