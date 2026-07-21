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
	_, err := buf.Read(outputBytes[:])
	if err != nil {
		t.Fatalf("failed to read from buffer: %s", err)
	}

	outputStr := string(bytes.Trim(outputBytes[:], "\x00"))
	if !strings.EqualFold(inputStr, outputStr) {
		t.Errorf("read frok buffer: get %q, expect %q", outputStr, inputStr)
	}

}
