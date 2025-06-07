package bubble

import (
	"testing"
)

func sorted(input []int) bool {
	if len(input) < 2 {
		return true
	}

	for i := range len(input) - 2 {
		if input[i] > input[i+1] {
			return false
		}
	}

	return true

}

func TestBubbleB(t *testing.T) {

	tests := []struct {
		input []int
	}{
		{[]int{7, 5, 8, 2, 1, 5, 4}},
		{[]int{5, 5, 4, 4, 3, 2, 2, 2, 11, 10}},
		{[]int{1, 11, 12, 2, 4, 14, 15, 5, 7, 17}},
	}

	for _, tc := range tests {
		Bubble(tc.input)
		if !sorted(tc.input) {
			t.Errorf("expect sorted slice, get %v", tc.input)
		}
	}

}
func TestBubble(t *testing.T) {

	tests := []struct {
		input []int
	}{
		{[]int{7, 5, 8, 2, 1, 5, 4}},
		{[]int{5, 5, 4, 4, 3, 2, 2, 2, 11, 10}},
		{[]int{1, 11, 12, 2, 4, 14, 15, 5, 7, 17}},
	}

	for _, tc := range tests {
		Sort(tc.input)
		if !sorted(tc.input) {
			t.Errorf("expect sorted slice, get %v", tc.input)
		}
	}

}

/*
Example:
{7, 5, 8, 2, 1, 5, 4}

i = 0
  j = 1, v = 7, w = 5: {5, 7, 8, 2, 1, 5, 4}
  j = 2, v = 5, w = 8:
  j = 3, v = 5, w = 2: {2, 7, 8, 5, 1, 5, 4}
  j = 4, v = 2, w = 1: {1, 7, 8, 5, 2, 5, 4}
  j = 5, v = 1, w = 5:
  j = 6, v = 1, w = 4:
i = 1
  j = 2, v = 7, w = 8:
  j = 3, v = 7, w = 5: {1, 5, 8, 7, 2, 5, 4}
  j = 4, v = 5, w = 2: {1, 2, 8, 7, 5, 5, 4}
  j = 5, v = 2, w = 5:
  j = 6, v = 2, w = 4:
 i = 2
   j = 3, v = 8, w = 7: {1, 2, 7, 8, 5, 5, 4}
   j = 4, v = 7, w = 5: {1, 2, 5, 8, 7, 5, 4}
   j = 5, v = 5, w = 5:
   j = 6, v = 5, w = 4: {1, 2, 4, 8, 7, 5, 5}
 i = 3
   j = 4, v = 8, w = 7: {1, 2, 4, 7, 8, 5, 5}
   j = 5, v = 7, w = 5: {1, 2, 4, 5, 8, 7, 5}
   j = 6, v = 5, w = 5:
 i = 4
   j = 5, v = 8, w = 7: {1, 2, 4, 5, 7, 8, 5}
   j = 6, v = 7, w = 5: {1, 2, 4, 5, 5, 8, 7}
 i = 5
   j = 6, v = 8, w = 7: {1, 2, 4, 5, 5, 7, 8}

*/
