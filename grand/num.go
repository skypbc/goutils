package grand

import (
	"crypto/rand"
	"math/big"

	"golang.org/x/exp/constraints"
)

func Num[T constraints.Integer](max T) (res T) {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return T(nBig.Uint64())
}

func Int64(max int64) int64 {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(max))
	return nBig.Int64()
}

func Uint64(max uint64) uint64 {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return nBig.Uint64()
}
