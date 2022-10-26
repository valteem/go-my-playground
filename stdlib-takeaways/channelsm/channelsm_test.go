package channelsm_test

import (
	. "lang.rev/stdlib-takeaways/channelsm"
	"fmt"
	"testing"
)

func TestMergeChannels(t *testing.T) {

	cb1 := make(chan string, 1)
	cb2 := make(chan string, 1)
	cb3 := make(chan string, 1)
	
// This works for buffered channels, no need for a goroutine as it is not blocked after sending to it
	cb1 <- "message 1"
	close(cb1)
	cb2 <- "message 2"
	close(cb2)
	cb3 <- "message 3"
	close(cb3)

	recb := MergeChannels(cb1, cb2, cb3)

	for c := range recb {
		fmt.Println(c)
	}

	cu1 := make(chan string)
	cu2 := make(chan string)
	cu3 := make(chan string)

	go func() {
		cu1 <- "message (u) 1"
		close(cu1)
		}()

	go func() {
		cu2 <- "message (u) 2"
		close(cu2)
		}()

	go func() {
		cu3 <- "message (u) 3"
		close(cu3)
		}()

	for c := range MergeChannels(cu1, cu2, cu3) {
		fmt.Println(c)
	}

}