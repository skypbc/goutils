package greader

import (
	"github.com/skypbc/goutils/gerrors"
	"io"
)

type ICacheReader interface {
	Read(buff []byte) (n int, err error)
	Read2(size int) (buff []byte, err error)

	Unread(buff []byte)
	Unread2(size int)

	Peek(offset int, buff []byte) (n int, err error)
	Peek2(offset int, size int) (buff []byte, err error)

	SetPosition(offset int) error
	Position() (int, error)
	ResetPosition()

	Data() []byte
	PopData() ([]byte, error)
	HardReset()

	Size() int
	MaxSize() int

	Close() error
}

type CacheReaderOpts struct {
	MinSize int
	MaxSize int
}

func NewCacheReader(r io.Reader, opts ...CacheReaderOpts) ICacheReader {
	minSize := 128
	maxSize := 0

	if len(opts) > 0 {
		if opts[0].MinSize > 0 {
			minSize = opts[0].MinSize
		}
		if opts[0].MaxSize > 0 {
			maxSize = opts[0].MaxSize
		}
	}

	return &cacheReader{
		source:  r,
		cache:   make([]byte, minSize),
		offset:  0,
		size:    0,
		maxSize: maxSize,
	}
}

type cacheReader struct {
	source  io.Reader
	cache   []byte
	offset  int
	size    int
	err     error
	maxSize int
}

func (c *cacheReader) Read(p []byte) (readed int, err error) {
	for {
		if c.size > 0 && c.offset < c.size {
			n := copy(p, c.cache[c.offset:c.size])
			c.offset += n
			readed += n
			return readed, nil
		}

		if c.err != nil {
			return readed, c.err
		}

		if err = c.extendCache(c.offset, len(p)); err != nil {
			return readed, err
		}

		n, err := c.source.Read(c.cache[c.size:])
		if err != nil {
			c.err = err
		}
		c.size += n
	}
}

func (c *cacheReader) Read2(size int) (buff []byte, err error) {
	for {
		if size <= c.size-c.offset {
			buff = c.cache[c.offset : c.offset+size]
			c.offset += size
			return buff, nil
		}

		if c.err != nil {
			return c.cache[c.offset:c.size], c.err
		}

		if err = c.extendCache(c.offset, size); err != nil {
			return c.cache[c.offset:c.size], c.err
		}

		n, err := c.source.Read(c.cache[c.size:])
		if err != nil {
			c.err = err
		}
		c.size += n
	}
}

func (c *cacheReader) Unread(buff []byte) {
	if len(buff) <= c.offset {
		for i := len(buff) - 1; i >= 0; i-- {
			c.offset--
			c.cache[c.offset] = buff[i]
		}
		return
	}

	buff = append(buff, c.cache[c.offset:c.size]...)
	c.cache = buff
	c.size = len(buff)
	c.offset = 0
}

func (c *cacheReader) Unread2(size int) {
	if c.offset -= size; c.offset < 0 {
		c.offset = 0
	}
}

func (c *cacheReader) Peek(offset int, buff []byte) (readed int, err error) {
	for {
		if c.size > 0 && offset < c.size {
			readed := copy(buff, c.cache[offset:c.size])
			return readed, nil
		}

		if c.err != nil {
			return 0, c.err
		}

		if err = c.extendCache(offset, len(buff)); err != nil {
			if offset < c.size {
				readed = copy(buff, c.cache[offset:c.size])
			}
			return readed, c.err
		}

		n, err := c.source.Read(c.cache[c.size:])
		if err != nil {
			c.err = err
		}
		c.size += n
	}
}

func (c *cacheReader) Peek2(offset int, size int) (buff []byte, err error) {
	for {
		if size <= c.size-offset {
			return c.cache[offset : offset+size], nil
		}

		if c.err != nil {
			return c.cache[offset:c.size], c.err
		}

		if err = c.extendCache(offset, size); err != nil {
			return c.cache[offset:c.size], c.err
		}

		n, err := c.source.Read(c.cache[c.size:])
		if err != nil {
			c.err = err
		}
		c.size += n
	}
}

func (c *cacheReader) SetPosition(offset int) error {
	if offset < 0 || offset > c.size {
		return gerrors.NewIncorrectParamsError()
	}
	c.offset = offset
	return nil
}

func (c *cacheReader) Position() (int, error) {
	return c.offset, c.err
}

func (c *cacheReader) ResetPosition() {
	c.offset = 0
}

func (c *cacheReader) PopData() (data []byte, err error) {
	data = c.Data()
	err = c.err
	c.HardReset()
	return data, err
}

func (c *cacheReader) HardReset() {
	c.size = 0
	c.offset = 0
	c.err = nil
}

func (c *cacheReader) Data() []byte {
	if c.size > 0 {
		return c.cache[c.offset:c.size]
	}
	return nil
}

func (c *cacheReader) Size() int {
	return c.size
}

func (c *cacheReader) MaxSize() int {
	return c.maxSize
}

func (c *cacheReader) Close() error {
	if closer, ok := c.source.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

func (c *cacheReader) extendCache(offset int, reqSize int) error {
	if reqSize > len(c.cache[offset:]) {
		if c.maxSize > 0 {
			if reqSize >= c.maxSize {
				return gerrors.NewShortBufferError().
					SetTemplate("Required size \"{req_size}\" exceeds max cache size \"{max_size}\"").
					AddInt("req_size", reqSize).
					AddInt("max_size", c.maxSize)

			}
			if len(c.cache)*2 >= c.maxSize {
				reqSize = c.maxSize
			} else {
				reqSize = len(c.cache) * 2
			}
		} else {
			if reqSize < len(c.cache)*2 {
				reqSize = len(c.cache) * 2
			}
		}
		tmp := make([]byte, reqSize)
		copy(tmp, c.cache)
		c.cache = tmp
	}
	return nil
}
