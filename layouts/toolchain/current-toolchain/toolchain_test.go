package main

import (
	"runtime"

	"testing"

	"toolchain/newer/newer"
)

func TestNewerToolchainversion(t *testing.T) {

	currentToolchainVersion := runtime.Version()
	newerToolchainVersion := newer.WrapToolchain()

	if newerToolchainVersion != currentToolchainVersion {
		t.Errorf("current and newer toolchain versions should be same:\nget %q and %q respectively",
			currentToolchainVersion,
			newerToolchainVersion)
	}
}
