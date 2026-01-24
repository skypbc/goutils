package greader

import (
	"errors"
	"io"
)

func NewCloser(r io.Reader, c ...io.Closer) io.ReadCloser {
	return &readerCloser{r, c}
}

type readerCloser struct {
	r io.Reader
	c []io.Closer
}

func (rc *readerCloser) Read(p []byte) (n int, err error) {
	return rc.r.Read(p)
}

func (rc *readerCloser) Close() (err error) {
	for _, c := range rc.c {
		if err2 := c.Close(); err2 != nil {
			if err != nil {
				err = errors.Join(err, err2)
			} else {
				err = err2
			}
		}
	}
	return err
}
