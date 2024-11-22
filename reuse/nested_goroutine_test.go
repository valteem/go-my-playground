package reuse_test

import (
	"testing"
)

func runNestedGoroutine(chIn, chOut chan bool) int {
	// nested goroutine
	go func(ch1, ch2 chan bool) {
		sg := <-ch1
		ch2 <- sg
	}(chIn, chOut)
	return 0
}

func TestNestedGoroutine(t *testing.T) {

	ch1, ch2 := make(chan bool), make(chan bool)

	runStatus := runNestedGoroutine(ch1, ch2)
	if runStatus != 0 {
		t.Errorf("run status: get %d, expect 0", runStatus)
	}

	ch1 <- true
	// Receiving from ch2 means that nested goroutine is still running
	// even after runNestedGoroutine() is completed
	transmitStatus := <-ch2
	if !transmitStatus {
		t.Errorf("failed to transmit status: %t", transmitStatus)
	}

}
