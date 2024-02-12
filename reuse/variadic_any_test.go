package reuse_test

import (
	"testing"

	"github.com/valteem/reuse"
)

func TestVariadicAny(t *testing.T) {
	var a1 string = "str"
	var a2 int = 1
	var a3 float32 = 1.0
	var a4 reuse.SliceOfBytes
	result := reuse.VarAny(a1, a2, a3, a4)
	expected := []string{"string", "int", "float32", "SliceOfBytes"}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("result %+v is not equal to %+v", v, expected[i])
		}
	}
}