package gbytes

import (
	crand "crypto/rand"
	"math/rand/v2"
)

func Random(size int) (res []byte) {
	if size == 0 {
		return nil
	}
	res = make([]byte, size)
	n, _ := crand.Read(res)
	return res[:n]
}

func Random2(size int, max ...int) (res []byte) {
	max_ := 0
	if len(max) > 0 && max[0] > 0 {
		max_ = max[0]
	}
	if size == 0 && max_ == 0 {
		return nil
	}
	if max_-size > 0 {
		size = rand.IntN(max_-size) + size
	}
	res = make([]byte, size)
	n, _ := crand.Read(res)
	return res[:n]
}
