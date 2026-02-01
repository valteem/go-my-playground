package reuse_test

import (
	"sync"
	"time"

	"testing"
)

var (
	chanCloseDelay = time.Millisecond * 500
)

func TestChanCloseDelay(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan int)

	go func() {
		for range c {
			<-c
		}
		wg.Done()
	}()

	now := time.Now()

	go func() {
		time.Sleep(chanCloseDelay)
		close(c)
	}()

	wg.Wait()
	after := time.Since(now)

	if after < chanCloseDelay {
		// Using %s format cause time.Duration has a built-in String() method
		t.Errorf("delay in main: get %s, expect no less than %s", after, chanCloseDelay)
	}
}
