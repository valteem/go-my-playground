package main

import (
	"reflect"

	"testing"
)

func TestLoopCapture(t *testing.T) {

	input := []string{"a", "b", "c"}

	out := make(chan string)

	for _, v := range input {
		go func() {
			out <- v
		}()
	}

	output := make([]string, 0)
	for range input {
		output = append(output, <-out)
	}

	expected := []string{"c", "c", "c"}

	// Test passes for 1.21 and fails for higher versions of Go
	// Newer toolchain (1.24.6) in go.mod still has no effect of pre-1.22 loop capture behaviour
	if !reflect.DeepEqual(expected, output) {
		t.Errorf("get %v, expect %v", output, expected)
	}

}
