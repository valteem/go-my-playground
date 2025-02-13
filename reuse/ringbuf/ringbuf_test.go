package ringbuf

import (
	"slices"
	"testing"
)

func TestAdd(t *testing.T) {

	rb := NewRingBuf[int]()

	for i := 0; i < 5; i++ {
		rb.Add(i)
	}

	bufExpected := []int{0, 1, 2, 3, 4, 0, 0, 0}
	if !slices.Equal(rb.Buf(), bufExpected) {
		t.Errorf("ring buffer content:\nget\n%v\nexpect\n%v\n", rb.Buf(), bufExpected)
	}

}
func TestRemainder(t *testing.T) {

	tests := []struct {
		idx    int
		size   int
		output int
	}{
		{0, 2, 1},
		{1, 2, 0},
		{2, 4, 3},
		{3, 4, 0},
	}

	for _, tc := range tests {
		output := (tc.idx + 1) % tc.size
		if output != tc.output {
			t.Errorf("%d %% %d: get %d, expect %d", tc.idx, tc.size, output, tc.output)
		}
	}

}
