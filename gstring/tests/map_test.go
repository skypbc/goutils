package tests

import (
	"github.com/skypbc/goutils/gfmt"
	"github.com/skypbc/goutils/gslice"
	"testing"
)

// go test -run TestAppend -v
// go test -bench . -run notest -benchtime=10000x

var intMap map[int]int
var stringMap map[string]int

var keySize = 100000
var steps = 1000
var intKeys []int
var stringKeys []string

func init() {
	for i := 0; i < keySize; i++ {
		intKeys = append(intKeys, i)
		stringKeys = append(stringKeys, gfmt.Sprintf("10.0.5.2.3.%d", i))
	}
	gslice.Shuffle(intKeys)
	gslice.Shuffle(stringKeys)
}

func BenchmarkCreateMapInt(b *testing.B) {
	intMap = map[int]int{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < keySize; i++ {
			intMap[intKeys[i]] = 1
		}
		b.StopTimer()
		intMap = map[int]int{}
		b.StartTimer()
	}
}

func BenchmarkCreateMapString(b *testing.B) {
	stringMap = map[string]int{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < keySize; i++ {
			stringMap[stringKeys[i]] = 1
		}
		b.StopTimer()
		stringMap = map[string]int{}
		b.StartTimer()
	}
}

func BenchmarkAccessMapInt(b *testing.B) {
	intMap = map[int]int{}
	for i := 0; i < keySize; i++ {
		intMap[intKeys[i]] = 1
	}
	b.ResetTimer()

	a := 0
	for i := 0; i < b.N; i++ {
		for i := 0; i < steps; i++ {
			a += intMap[intKeys[i]]
		}
	}
}

func BenchmarkAccessMapString(b *testing.B) {
	stringMap = map[string]int{}
	for i := 0; i < keySize; i++ {
		stringMap[stringKeys[i]] = 1
	}
	b.ResetTimer()

	a := 0
	for i := 0; i < b.N; i++ {
		for i := 0; i < steps; i++ {
			a += stringMap[stringKeys[i]]
		}
	}
}
