package gmap

import (
	"encoding/json"
)

func Clone[M ~map[K]V, K comparable, V any](m M) (M, error) {
	res := make(M, len(m))
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func Copy[M ~map[K]V, K comparable, V any](m M) M {
	if m == nil {
		return nil
	}
	res := make(M, len(m))
	for k, v := range m {
		res[k] = v
	}
	return res
}
