package reuse_test

import (
	"fmt"
	"testing"
)

func TestSliceCopy(t *testing.T) {

	source := []int{1, 2, 3}

	copy := source // `copy` and `source` point to the same memory block
	copy[0] = 0
	if source[0] != 0 {
		t.Errorf("expect shallow copy")
	}

	sourcePointer, copyPointer := fmt.Sprintf("%p", &source[0]), fmt.Sprintf("%p", &copy[0])
	if sourcePointer != copyPointer {
		t.Errorf("expect same memory addresses, get %s, %s", sourcePointer, copyPointer)
	}

	source = make([]int, 1, 16) // `source` points to another memory block
	if copy[0] != 0 || copy[1] != 2 || copy[2] != 3 {
		t.Errorf("copy should stay unchanged, get %v instead", copy)
	}

	sourcePointer, copyPointer = fmt.Sprintf("%p", &source[0]), fmt.Sprintf("%p", &copy[0])
	if sourcePointer == copyPointer {
		t.Errorf("expect different memory addresses, get %s, %s", sourcePointer, copyPointer)
	}

}
