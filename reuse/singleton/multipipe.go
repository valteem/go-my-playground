package singleton

import (
	"fmt"
)

type MultiPipe struct {
	numInputs int
	output    chan string
}

func (mp MultiPipe) Validate() error {
	if mp.numInputs < 1 {
		return fmt.Errorf("expect at least one input, get %d", mp.numInputs)
	}
	return nil
}

func (mp MultiPipe) Read() string {
	return <-mp.output
}

// TODO: add input(s) feeding
