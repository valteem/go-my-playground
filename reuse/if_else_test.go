package reuse_test

import (
	"testing"
)

const (
	x = 5
	y = 4
)

func bmCmpIfElse(x, y int) int {
	diff := x - y
	counter := 0
	if diff < 0 {
		diff = 0
	} else {
		for i := 0; i < diff; i++ {
			counter++
		}
	}
	return counter
}

func bmCmpIfOnly(x, y int) int {
	diff := x - y
	counter := 0
	if diff < 0 {
		diff = 0
	}
	for i := 0; i < diff; i++ {
		counter++
	}
	return counter
}

func BenchmarkIfElsevsIfOnly(b *testing.B) {

	b.Run("IfElse", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bmCmpIfElse(x, y)
		}
	})

	b.Run("IfOnly", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bmCmpIfOnly(x, y)
		}
	})

}
