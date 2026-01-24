package gmap

import "encoding/json"

func Size[M ~map[K]V, K comparable, V any](m M) (int64, error) {
	if m == nil {
		return 0, nil
	}
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return 0, err
	}
	return int64(len(jsonBytes)), nil
}
