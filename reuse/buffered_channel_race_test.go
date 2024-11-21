package reuse_test

import (
	"sync"
	"testing"
)

const (
	numWriters       = 1000
	numWrites        = 1000
	numReadsExpected = numWriters * numWrites
)

// Buffered channels are not subject to data race and can be used instead of mutexes in certain scanarios
func TestBufferedChannelRace(t *testing.T) {

	ch := make(chan int, numWriters)

	var numReads int

	for i := 0; i < numWriters; i++ {
		go func() {
			for i := 0; i < numWrites; i++ {
				ch <- 1
			}

		}()
	}

	var wg sync.WaitGroup
	wg.Add(numReadsExpected)

	go func() {
		for range ch {
			numReads++
			wg.Done()
		}
	}()

	wg.Wait()

	if numReads != numReadsExpected {
		t.Errorf("get %d, expect %d", numReads, numReadsExpected)
	}

}
