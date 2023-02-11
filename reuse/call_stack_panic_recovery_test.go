package reuse_test

import (
	"testing"
	reuse "github.com/valteem/reuse"
)

func TestCallStackPanicREcover(t *testing.T) {
	reuse.CallStackPanicRecover()
}
