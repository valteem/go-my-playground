package reuse_test

import (
	"fmt"
	"testing"
	"github.com/valteem/reuse"
)



func TestNextPowerOf2(t *testing.T) {
	test := []uint32{3, 4, 6, 8, 12, 16}
	for _, i := range test {
		fmt.Println(i, reuse.NextPowerOf2(uint32(i)))
	}
}