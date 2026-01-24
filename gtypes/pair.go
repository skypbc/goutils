package gtypes

type PairItem[V1, V2 any] struct {
	Value1 V1
	Value2 V2
}

type PairList[V1, V2 any] []PairItem[V1, V2]

type MapItem[K, V any] struct {
	Key   K
	Value V
}

type MapItemList[K, V any] []MapItem[K, V]
