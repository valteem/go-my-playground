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

func sumOfSlice(input []int) int {
	input = append(input, 0) // replaces input[4] for 0 if invoked as sumOfSlice(input[0:4])
	s := 0
	for input[0] > 0 {
		s += input[0]
		input = input[1:]
	}
	return s
}
func TestAppendInFunction(t *testing.T) {

	input := []int{1, 1, 1, 1, 1}

	if outputActual, outputExpected := sumOfSlice(input[0:4]), 4; outputActual != outputExpected {
		t.Errorf("[0:4]: get %d, expect %d", outputActual, outputExpected)
	}
	// input: [1, 1, 1, 1, 0]

	if outputActual, outputExpected := sumOfSlice(input[2:5]), 2; outputActual != outputExpected {
		t.Errorf("[0:4]: get %d, expect %d", outputActual, outputExpected)
	}

}

func TestSeqAppend(t *testing.T) {

	t.Run("append on pointer to slice", func(t *testing.T) {
		s := make([]int, 0)
		f := func(s *[]int, i int) {
			*s = append(*s, i)
		}
		for i := range 16 {
			f(&s, i)
			fmt.Println(s, len(s), cap(s))
		}
	})

	t.Run("append on slice", func(t *testing.T) {
		s := make([]int, 0)
		f := func(s []int, i int) {
			s = append(s, i)
		}
		for i := range 16 {
			f(s, i)
			fmt.Println(s, len(s), cap(s))
		}
	})

}

func TestCopySlice(t *testing.T) {

	a := []int{0, 1, 2, 3}

	var b []int
	count := copy(b, a)
	if count > 0 {
		t.Errorf("expect 0 elements copied, get %d", count)
	}

	c := make([]int, 3)
	count = copy(c, a)
	if count != 3 {
		t.Errorf("expect %d elements copied, get %d", len(c), count)
	}

}

func TestSliceRange(t *testing.T) {

	type Food struct {
		name    string
		expired bool
	}

	t.Run("slice of struct values", func(t *testing.T) {
		f := []Food{{"apples", false}, {"cherries", false}, {"onions", false}}

		for _, foodItem := range f {
			foodItem.expired = true
		}

		for _, foodItem := range f {
			if foodItem.expired {
				t.Errorf("%q are not expected to expire", foodItem.name)
			}
		}
	})

	t.Run("slice of struct pointers", func(t *testing.T) {
		f := []*Food{{"apples", false}, {"cherries", false}, {"onions", false}}

		for _, foodItem := range f {
			foodItem.expired = true
		}

		for _, foodItem := range f {
			if !foodItem.expired {
				t.Errorf("%q are expected to expire", foodItem.name)
			}
		}
	})

}
