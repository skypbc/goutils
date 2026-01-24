package gmap

import (
	"bytes"
	"encoding/json"
)

func EqualJson[M1 ~map[K]V, M2 ~map[K]V, K comparable, V comparable](m1 M1, m2 M2) (bool, error) {
	if len(m1) != len(m2) {
		return false, nil
	}
	b1, err := json.Marshal(m1)
	if err != nil {
		return false, err
	}
	b2, err := json.Marshal(m2)
	if err != nil {
		return false, err
	}
	if len(b1) != len(b2) {
		return false, nil
	}
	return bytes.Equal(b1, b2), nil
}
