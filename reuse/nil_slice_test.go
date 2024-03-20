// https://stackoverflow.com/a/44305910

package reuse_test

import (
	"testing"
)

func TestInitNilSlice(t *testing.T) {

	in := []string(nil) // most probably useful for inline nil slice declaration
	if in != nil {
		t.Errorf("slice must be nil")
	}
	add := []string{"first, second", "third", "closing"}
	out := append(in, add...)
	for i, v := range out {
		if v != add[i] {
			t.Errorf("two elements must be equal: %v and %v", v, add[i])
		}
	}
}
