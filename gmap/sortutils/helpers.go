package sortutils

import (
	"fmt"
	"github.com/skypbc/goutils/gtypes"
)

func DebugPrintMapItemList[K, V any](pairs gtypes.MapItemList[K, V]) {
	fmt.Printf("%T:\n", pairs)
	for i, pair := range pairs {
		fmt.Printf("%2s- Index: %d, Key: %+v, Value: %+v\n", "", i, pair.Key, pair.Value)
	}
}
