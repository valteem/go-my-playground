package reuse_test

import (
	"testing"
)

func TestChannelRead(t *testing.T) {

	c := make(chan int, 5)

	for i := 0; i < 5; i++ {
		c <- i
	}

	for i := 0; i < 5; i++ {
		j, ok := <-c
		if j != i || !ok {
			t.Errorf("reading from buffered channel: get %d, %t, expect %d, %t", j, ok, i, true)
		}
	}

	close(c)

	for i := 0; i < 5; i++ {
		j, ok := <-c
		if j != 0 || ok {
			t.Errorf("reading from closed channel: get %d, %t, expect %d, %t", j, ok, 0, false)
		}
	}

}
