package benchmarkexample

import (
	"math/rand"
	"testing"
)

const (
	sliceLength = 100000
	maxValue = 100
)

var (
	result int // https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
)

func setSlice() []int {

	s := make([]int, sliceLength)

	for i := 0; i < sliceLength; i++ {
		s[i] = int(rand.Float64() * maxValue)
	}

	return s

}

func BenchmarkLoopRange(b *testing.B) {

	var r int

	s:= setSlice()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r = LoopRange(s)
	}

	result = r

}

func BenchmarkLoopFor(b *testing.B) {

	var r int

	s:= setSlice()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r = LoopFor(s)
	}

	result = r

}