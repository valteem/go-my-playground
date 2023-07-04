package reuse_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/valteem/reuse"
)

func TestErrorUnwrap(t *testing.T) {
	err := errors.Unwrap(reuse.Outer(-1))
	fmt.Println(err)
	if errors.Is(err, reuse.ErrNotNegativeValue) {
		fmt.Println("sentinel error found")
	}
}