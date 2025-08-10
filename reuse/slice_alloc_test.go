package reuse_test

import (
	"bytes"

	"testing"
)

func TestSliceAlloc(t *testing.T) {

	// https://github.com/prometheus/prometheus/blob/c70163603458822ee68ac2bf8f66899dad57c359/model/labels/labels_common.go#L46
	// On stack to avoid memory allocation while building the output
	var bytea [1024]byte
	b := bytes.NewBuffer(bytea[:0])

	input := "some string"
	n, err := b.Write([]byte(input))
	if err != nil {
		t.Fatalf("failed to write to buffer: %v", err)
	}

	underlying := string(bytea[:n])
	if underlying != input {
		t.Errorf("underlying array sliced to %d bytes: get %q, expect %q", n, underlying, input)
	}

}

func TestAllocated(t *testing.T) {
	s := make([]int32, 0, 1<<25)
	for i := range 1 << 25 {
		s = append(s, int32(i))
		if actual, expected := len(s), i+1; actual != expected {
			t.Errorf("step %d: get slice length %d, expect %d", i, actual, expected)
		}
		if actual, expected := cap(s), 1<<25; actual != expected {
			t.Errorf("step %d: get slice capacity %d, expect %d", i, actual, expected)
		}
	}
}
