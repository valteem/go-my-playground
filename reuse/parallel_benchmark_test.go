package reuse_test

import (
	"testing"
)

func BenchmarkParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var s float32
			for i := 1; i <= 10000; i++ {
				s = (s + float32(i)) / float32(i)
			}	
		}
	})
}