package singleton

import (
	"testing"
)

func TestProperlySetPipe(t *testing.T) {

	mp := &MultiPipe{numInputs: 1, output: make(chan string, 1)} // use buffered channel to avoid blocking

	SetGlobalPipe(mp)

	if Validate() != nil {
		t.Errorf("failed to validate properly set MultiPipe: get %s", Validate())
	}

	input := "some input"
	mp.output <- "some input"

	if output := Read(); output != input {
		t.Errorf("output: get %q, expect %q", output, input)
	}
}

func TestPipeValidatePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expect panic after calling Validate on globalPipe that is not set")
		}
	}()
	Validate()
}

func TestPipeReadPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expect panic after calling Read of gllobalPipe that is not set")
		}
	}()
	Read()
}
