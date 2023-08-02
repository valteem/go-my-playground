package reuse_test

import (
	"testing"

	"github.com/valteem/reuse"
)

func TestGuardInit(t *testing.T) {

	reuse.Init()

	reuse.Init()

	reuse.Stop()

}

func TestGuardStop(t *testing.T) {

	reuse.Init()

	reuse.Stop()

	reuse.Stop()

}