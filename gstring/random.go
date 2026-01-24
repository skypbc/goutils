package gstring

import (
	"github.com/skypbc/goutils/gnum"
)

func RandomString(alphabet string, width int) string {
	return string(RandomRune([]rune(alphabet), width))
}

func RandomRune(alphabet []rune, width int) []rune {
	max := int64(len(alphabet))
	buff := make([]rune, 0, max)
	for i := 0; i < width; i++ {
		buff = append(buff, alphabet[gnum.RandInt64(max)])
	}
	return buff
}
