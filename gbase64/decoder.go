package gbase64

import (
	"encoding/base64"
	"io"
)

type decoder struct {
	r io.Reader
}

func NewDecoder(enc *base64.Encoding, r io.Reader) io.Reader {
	return &decoder{r: base64.NewDecoder(enc, r)}
}

func (d *decoder) Read(buff []byte) (n int, err error) {
	return d.r.Read(buff)
}
