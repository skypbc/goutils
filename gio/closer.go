package gio

import (
	"github.com/skypbc/goutils/gerrors"
	"io"
)

func NewCloser(items ...io.Closer) *closer {
	return &closer{
		items: items,
	}
}

type closer struct {
	items []io.Closer
}

func (c *closer) Close() (err error) {
	if len(c.items) == 0 {
		return nil
	}
	for index := range c.items {
		closer := c.items[index]
		if closer == nil {
			continue
		}
		if err2 := closer.Close(); err2 != nil {
			if err != nil {
				err = gerrors.Wrap(err, err2)
			} else {
				err = gerrors.Wrap(err2)
			}
		}
		c.items[index] = nil
	}
	c.items = nil
	return err
}

func (c *closer) Append(items ...io.Closer) {
	for _, item := range items {
		if item == nil {
			continue
		}
		c.items = append(c.items, item)
	}
}

func (c *closer) Count() int {
	return len(c.items)
}
