package greader

import (
	"bytes"
	"io"
)

type Unreder interface {
	Unread(b []byte)
}

func Join(cacheReader ICacheReader, stream io.Reader) io.Reader {
	raw, _ := cacheReader.PopData()
	return Join2(raw, stream)
}

func Join2(raw []byte, stream io.Reader) io.Reader {
	switch x := stream.(type) {
	case Unreder:
		x.Unread(raw)
	case io.Closer:
		stream = NewCloser(io.MultiReader(bytes.NewReader(raw), stream), x)
	default:
		stream = io.MultiReader(bytes.NewReader(raw), stream)
	}
	return stream
}
