package reuse_test

import (
	"sync"
	"testing"
)

const (
	numWriters = 100
	numWrites  = 10000
	bufSize    = numWriters * numWrites
)

// Buffered channels are not subject to data race and can be used instead of mutexes in certain scanarios
func TestBufferedChannelRace(t *testing.T) {

	ch := make(chan int, bufSize)

	var numReads int

	for i := 0; i < numWriters; i++ {
		go func() {
			for i := 0; i < numWrites; i++ {
				ch <- 1
			}

		}()
	}

	var wg sync.WaitGroup
	wg.Add(bufSize)

	go func() {
		for range ch {
			numReads++
			wg.Done()
		}
	}()

	wg.Wait()

	if numReads != bufSize {
		t.Errorf("get %d, expect %d", numReads, bufSize)
	}

}
