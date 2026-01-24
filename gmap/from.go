package gmap

import "encoding/json"

func From[K comparable, V any](value any) (res map[K]V, err error) {
	data, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &res)
	return res, err
}

func FromSlice[T any, K comparable, V any](s []T, f func(i int) (K, V)) (res map[K]V) {
	m := make(map[K]V, len(s))
	for i := range s {
		k, v := f(i)
		m[k] = v
	}
	return m
}
