package gio

import (
	"errors"
	"io"
)

type MultiReaderCloser struct {
	closers []io.Closer
	reader  io.Reader
}

func NewMultiReaderCloser(readers ...io.Reader) *MultiReaderCloser {
	rc := &MultiReaderCloser{}
	for _, r := range readers {
		if c, ok := r.(io.Closer); ok {
			rc.closers = append(rc.closers, c)
			continue
		}
	}
	rc.reader = io.MultiReader(readers...)
	return rc
}

func (m *MultiReaderCloser) Read(p []byte) (n int, err error) {
	return m.reader.Read(p)
}

func (m *MultiReaderCloser) Close() (err error) {
	for _, c := range m.closers {
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
