package main

import (
	"runtime"

	"testing"
)

func TestAnything(t *testing.T) {
	if actual, expected := runtime.Version(), "go1.24.5"; actual != expected {
		t.Errorf("toolchain vaersion: get %q, expect %q", actual, expected)
	}
}
