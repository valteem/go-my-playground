package main

import (
	"testing"

	"toolchain/newer/newer"
)

func TestNewerToolchainversion(t *testing.T) {
	// Get version set in go.mod in module version, if toolchain is not set
	if actual, expected := newer.WrapToolchain(), "go1.23.2"; actual != expected {
		t.Errorf("toolchain version: get %q, expect %q", actual, expected)
	}
}
