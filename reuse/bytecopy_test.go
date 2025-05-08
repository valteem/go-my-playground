package reuse_test

import (
	"testing"
)

func copyBytes(src []byte) []byte {
	if len(src) == 0 {
		return nil
	}
	dst := make([]byte, len(src))
	copy(dst, src)
	return dst
}

func TestCopyBytes(t *testing.T) {

	src := []byte{1, 1, 1}
	dst := copyBytes(src)

	for i, b := range src {
		src[i] = byte(2 * int(b))
	}

	for i, b := range dst {
		if v := int(b); v != 1 {
			t.Errorf("dts[%d]: get %d, expect %d", i, v, 1)
		}
	}

}

func BenchmarkCopyBytes(b *testing.B) {

	src := []byte{}
	for i := range 128 {
		src = append(src, byte(i))
	}

	b.Run("copyBytes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copyBytes(src)
		}
	})

	b.Run("append(nil, ...)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = append([]byte(nil), src...)
		}
	})

	b.Run("append({}, ...)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = append([]byte{}, src...)
		}
	})

}
