package grand

import (
	"github.com/skypbc/goutils/gslice"
)

func Shuffle[S ~[]T, T any](s S) {
	gslice.Shuffle[S](s)
}
