package reuse_test

import (
	"fmt"
	"testing"

	"github.com/valteem/reuse"
)

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go reuse.WriteChan(ch1, []int{1, 2, 3})
	go reuse.WriteChan(ch2, []int{11, 12, 13})
	go reuse.WriteChan(ch3, []int{21, 22, 23})
	mch := reuse.MergeChannels(ch1, ch2, ch3)
	for v := range mch {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")		
}