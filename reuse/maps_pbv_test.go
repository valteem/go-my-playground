// Maps, like channels, but unlike slices, are just pointers to runtime types.
// A map is just a pointer to a runtime.hmap structure.

package reuse_test

import (
	"testing"
)

func mapPassByValue(m map[string]string) {
	m["cloud"] = "white"
}
func TestMapsPassByWhat(t *testing.T) {

	m := map[string]string{"fruit": "apple", "veg": "onion"}

	mapPassByValue(m)

	if m["cloud"] != "white" {
		t.Errorf("failed to add new map entry")
	}

}
