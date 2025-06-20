package reuse_test

import (
	"unsafe"

	"testing"
)

func yoloString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func BenchmarkByteToString(b *testing.B) {

	output := "some string of moderate length but long enough to produce meaningful results"
	input := []byte(output)

	b.Run("standard conversion", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = string(input)
		}
	})

	b.Run("yoloString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = yoloString(input)
		}
	})

}
