package reuse_test

import (
	"testing"
)

type SmallObject struct {
	a, b, c, d int
}
type PlainObject struct {
	a, b, c, d int
	e          int
}

func sum1(obj SmallObject) int {
	return obj.a + obj.b + obj.c + obj.d
}
func sum2(obj PlainObject) int {
	return obj.a + obj.b + obj.c + obj.d + obj.e
}

func BenchmarkSmallAndPlainObjects(b *testing.B) {
	b.Run("SmallObject", func(b *testing.B) {
		var ret int
		for i := 0; i < b.N; i++ {
			obj := SmallObject{a: i, d: i}
			ret = sum1(obj)
		}
		_ = ret
	})
	b.Run("PlainObject", func(b *testing.B) {
		var ret int
		for i := 0; i < b.N; i++ {
			obj := PlainObject{a: i, d: i, e: i}
			ret = sum2(obj)
		}
		_ = ret
	})
}
