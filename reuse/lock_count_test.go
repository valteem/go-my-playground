package reuse_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

var (
	lockedCounter              = 0
	countLock                  sync.Mutex
	atomicCounter              atomic.Int64
	numWorkers                 = 1000
	numCallsToIncrementCounter = 1000
)

func TestLockedCounter(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(wg *sync.WaitGroup) {
			for i := 0; i < numCallsToIncrementCounter; i++ {
				countLock.Lock()
				lockedCounter++
				countLock.Unlock()
			}
			wg.Done()
		}(&wg)

	}

	wg.Wait()

	if lockedCounter != numWorkers*numCallsToIncrementCounter {
		t.Errorf("counter: get %d, expect %d", lockedCounter, numWorkers*numCallsToIncrementCounter)
	}

}

// Atomics faster than mutexes
// https://stackoverflow.com/a/47446114
func TestAtomicCounter(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(wg *sync.WaitGroup) {
			for i := 0; i < numCallsToIncrementCounter; i++ {
				atomicCounter.Add(1)
			}
			wg.Done()
		}(&wg)
	}

	wg.Wait()

	if actual, expected := atomicCounter.Load(), int64(numWorkers*numCallsToIncrementCounter); actual != expected {
		t.Errorf("counter value:\nget\n%d\nexpect\n%d", actual, expected)
	}

}

// escape analysis
// go test -gcflags "-m" <filename>
func BenchmarkCounter(b *testing.B) {

	b.Run("mutex", func(b *testing.B) {
		var wg sync.WaitGroup
		wg.Add(b.N)
		for i := 0; i < b.N; i++ {
			go func(wg *sync.WaitGroup) {
				for i := 0; i < numCallsToIncrementCounter; i++ {
					countLock.Lock()
					lockedCounter++
					countLock.Unlock()
				}
				wg.Done()
			}(&wg)
		}
		wg.Wait()
	})

	b.Run("atomic", func(b *testing.B) {
		var wg sync.WaitGroup
		wg.Add(b.N)
		for i := 0; i < b.N; i++ {
			go func(wg *sync.WaitGroup) {
				for i := 0; i < numCallsToIncrementCounter; i++ {
					atomicCounter.Add(1)
				}
				wg.Done()
			}(&wg)
		}
		wg.Wait()
	})
}

func BenchmarkCounterNoConcurrency(b *testing.B) {

	b.Run("mutex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			countLock.Lock()
			lockedCounter++
			countLock.Unlock()
		}
	})

	b.Run("atomic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			atomicCounter.Add(1)
		}
	})

}
