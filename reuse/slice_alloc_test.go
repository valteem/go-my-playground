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
