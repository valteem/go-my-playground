package reuse_test

import (
	"sync"
	"testing"
	"time"
)

var (
	lockedCounter              = 0
	countLock                  sync.Mutex
	numCallsToIncrementCounter = 1000
)

func incrementLockedCounter(timeout time.Duration, wg *sync.WaitGroup) {
	countLock.Lock()
	lockedCounter++
	countLock.Unlock()
	time.Sleep(timeout)
	wg.Done()
}

func TestLockedCounter(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(numCallsToIncrementCounter)

	timeout := 5 * time.Second
	for i := 0; i < numCallsToIncrementCounter; i++ {
		go incrementLockedCounter(timeout, &wg)
	}

	wg.Wait()

	if lockedCounter != numCallsToIncrementCounter {
		t.Errorf("counter: get %d, expect %d", lockedCounter, numCallsToIncrementCounter)
	}

}
