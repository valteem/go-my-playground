package reuse_test

import (
	"testing"
)

var (
	inputString   = "input string"
	anotherString = "another string"
)

func BenchmarkStringReplace(b *testing.B) {
	b.Run("ByValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = func(s string) string {
				return s + " " + anotherString
			}(inputString)
		}
	})
	b.Run("ByPointer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = func(s *string) string {
				return *s + " " + anotherString
			}(&inputString)
		}
	})
}
