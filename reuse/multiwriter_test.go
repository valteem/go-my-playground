package reuse_test

import (
	"bytes"
	"testing"

	"github.com/valteem/reuse"
)

func TestMultiWriter(t *testing.T) {

	inputString := []string{"one", " ", "ring", " ", "to", " ", "rule", " ", "them", " ", "all"}

	mw := reuse.MultiWriter{}

	w1, w2 := bytes.Buffer{}, bytes.Buffer{}
	// https://stackoverflow.com/a/53263087
	// Because the way bytes.Buffer implements io.Writer is
	/*
	   func (b *Buffer) Write(p []byte) (n int, err error) {
	       ...
	   }
	*/
	mw.AddWriter(&w1)
	mw.AddWriter(&w2)

	for _, input := range inputString {
		mw.Write([]byte(input))
	}

	outputExpected := "one ring to rule them all"
	o1 := w1.String()
	if o1 != outputExpected {
		t.Errorf("writer #1: get %s, expect %s", o1, outputExpected)
	}
	o2 := w2.String()
	if o2 != outputExpected {
		t.Errorf("writer #2: get %s, expect %s", o2, outputExpected)
	}
}
