package reuse_test

import (
	"testing"
)

func BenchmarkForVsIf(b *testing.B) {

	b.Run("For", func(b *testing.B) {
		count := 0
		for i := 0; i < b.N; i++ {
			for (i/2)*2 == i {
				count++
				break
			}
		}
	})

	b.Run("If", func(b *testing.B) {
		count := 0
		for i := 0; i < b.N; i++ {
			if (i/2)*2 == i {
				count++
			}
		}

	})

}
