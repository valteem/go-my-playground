// go test -bench=. -benchmem -memprofile isolated.out
// go tool pprof isolated.out
// list <TestName>

package isolated

import (
	"sync"
	"sync/atomic"
	"testing"
)

const (
	numCalls = 1000
)

var (
	wg sync.WaitGroup
)

func BenchmarkCounter(b *testing.B) {
	var counterLock sync.Mutex
	var counter int
	var atomicCounter atomic.Int64

	b.Run("mutex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				for i := 0; i < numCalls; i++ {
					counterLock.Lock()
					counter++
					counterLock.Unlock()
				}
				wg.Done()
			}(&wg)
			wg.Wait()
		}
	})

	b.Run("atomic", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				for i := 0; i < numCalls; i++ {
					atomicCounter.Add(1)
				}
				wg.Done()
			}(&wg)
			wg.Wait()
		}
	})
}

// Using wg.Add() and wg.wait inside benchmark iteration eliminates lockSlow()/unlockSlow() calls
// resulting from thousands of goroitins trying to access single lock (almost) at the same time
// Removing waitgroup as goroutine parameter eliminates another allocation for both benchmarks
// Hence this version of benchmark produces 1 alloc/op for both "mutex" and "atomic"
func BenchmarkCounterGlobalWaitGroup(b *testing.B) {
	var counterLock sync.Mutex
	var counter int
	var atomicCounter atomic.Int64

	b.Run("mutex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < numCalls; i++ {
					counterLock.Lock()
					counter++
					counterLock.Unlock()
				}
				wg.Done()
			}()
			wg.Wait()
		}
	})

	b.Run("atomic", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < numCalls; i++ {
					atomicCounter.Add(1)
				}
				wg.Done()
			}()
			wg.Wait()
		}
	})

}
