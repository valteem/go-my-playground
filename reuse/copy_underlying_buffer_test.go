package reuse_test

import (
	"fmt"
	"testing"

	"github.com/valteem/reuse"
)

func TestCopyUnderlyingBuffer(t *testing.T) {
	fmt.Println(reuse.CopyUnderlyingBuffer("text"))
}