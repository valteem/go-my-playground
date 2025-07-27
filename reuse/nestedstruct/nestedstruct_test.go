package nestedstruct

import (
	"testing"
)

func TestPromotedShadowed(t *testing.T) {

	c1 := C1{A: A{}, b: B{}}
	if actual, expected := c1.Write(), "A"; actual != expected {
		t.Errorf("promoting A.Write(): get %q, expect %q", actual, expected)
	}

	c2 := C2{a: A{}, B: B{}}
	if actual, expected := c2.Write(), "B"; actual != expected {
		t.Errorf("promoting B.Write(): get %q, expect %q", actual, expected)
	}

	c1b := C1B{A: A{}, b: B{}}
	if actual, expected := c1b.Write(), "B"; actual != expected {
		t.Errorf("promoting A.Write(): get %q, expect %q", actual, expected)
	}

	c2a := C2A{a: A{}, B: B{}}
	if actual, expected := c2a.Write(), "A"; actual != expected {
		t.Errorf("promoting B.Write(): get %q, expect %q", actual, expected)
	}
}
