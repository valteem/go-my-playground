package reuse_test

import (
	"bytes"
	"strings"

	"testing"
)

func TestByteBufTrim(t *testing.T) {

	var bytea [1024]byte
	buf := bytes.NewBuffer(bytea[:0])

	inputStr := "one apple is better than a bag of potatoes"
	inputByte := []byte(inputStr)

	for _, b := range inputByte {
		buf.WriteByte(b)
	}

	var outputBytes [1024]byte
	// proper way to read buffer content - initial slice is not touched
	_, err := buf.Read(outputBytes[:])
	if err != nil {
		t.Fatalf("failed to read from buffer: %s", err)
	}

	outputStr := string(bytes.Trim(outputBytes[:], "\x00"))
	if !strings.EqualFold(inputStr, outputStr) {
		t.Errorf("read frok buffer: get %q, expect %q", outputStr, inputStr)
	}

}

func TestBufMultipleWrite(t *testing.T) {

	var bytea [1024]byte

	// Use of undelying slice in the bytes buffer is completely counterintuitive:
	// one may suppose that using a slice as NewBuffer argument prepares it for future writes.
	// However, the present content of the slice is intended for future reads (not writes),
	// and any new write will add data beyond the initial length of the underlying slice
	b1 := bytes.NewBuffer(bytea[:0])
	b2 := bytes.NewBuffer(bytea[:7])
	b3 := bytes.NewBuffer(bytea[16:]) // does not add anything

	_, _ = b1.WriteString("apples")
	_, _ = b2.WriteString("potatoes")
	_, _ = b3.WriteString("flowers")

	// Direct manipulations with the slice behind bytes buffer should be avoided in real life
	// As per documentation:
	// The caller should not use (initial slice) after (NewBuffer) call
	output, expected := string(bytes.Trim(bytea[:], "\x00")), "apples\x00potatoes"
	if output != expected {
		t.Errorf("get %q, expect %q", output, expected)
	}
}

func TestSlicedArrayCap(t *testing.T) {

	var bytea [1024]byte
	b := bytea[:0] // zero-length slice with capacity equal to initial array size

	lenActual, lenExpected, capActual, capExpected := len(b), 0, cap(b), 1024
	if lenActual != lenExpected {
		t.Errorf("sliced array length: get %d, expect %d", lenActual, lenExpected)
	}
	if capActual != capExpected {
		t.Errorf("sliced array capacity: get %d, expect %d", capActual, capExpected)
	}

}
