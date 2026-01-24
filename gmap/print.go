package gmap

import (
	"encoding/json"
	"github.com/skypbc/goutils/gfmt"
)

type PrintOpts struct {
	Indent *string
	Prefix *string
}

func Print[M ~map[K]V, K comparable, V any](m M, opts ...PrintOpts) error {
	text, err := TrySprint(m, opts...)
	if err != nil {
		return err
	}
	_, err = gfmt.Print(text)
	return err
}

func Sprint[M ~map[K]V, K comparable, V any](m M, opts ...PrintOpts) string {
	text, err := TrySprint(m, opts...)
	if err != nil {
		return ""
	}
	return text
}

func TrySprint[M ~map[K]V, K comparable, V any](m M, opts ...PrintOpts) (string, error) {
	ident := "    "
	prefix := ""

	if len(opts) > 0 {
		if opts[0].Indent != nil {
			ident = *opts[0].Indent
		}
		if opts[0].Prefix != nil {
			prefix = *opts[0].Prefix
		}
	}

	data, err := json.MarshalIndent(m, prefix, ident)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
