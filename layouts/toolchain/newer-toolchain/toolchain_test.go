package main

import (
	"runtime"

	"testing"
)

func TestAnything(t *testing.T) {
	if actual, expected := runtime.Version(), "go1.25.5"; actual != expected {
		t.Errorf("toolchain version: get %q, expect %q", actual, expected)
	}
}
