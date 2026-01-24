package gnum

import (
	"crypto/rand"
	"math/big"

	"golang.org/x/exp/constraints"
)

func Rand[T constraints.Integer](max T) (res T) {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return T(nBig.Uint64())
}

func RandInt64(max int64) int64 {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(max))
	return nBig.Int64()
}

func RandUint64(max uint64) uint64 {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return nBig.Uint64()
}
