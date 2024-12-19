// go test -bench=. -benchmem -memprofile counter.out
// go tool pprof counter.out
// list <TestName>

package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

const (
	numCalls = 1000
)

var (
	globalCounterLock   sync.Mutex
	globalCounter       int
	globalAtomicCounter atomic.Int64
	wg                  sync.WaitGroup
)

func BenchmarkCounter(b *testing.B) {
	var counterLock sync.Mutex
	var counter int
	var atomicCounter atomic.Int64

	b.Run("mutex", func(b *testing.B) {
		wg.Add(b.N)
		for i := 0; i < b.N; i++ {
			go func(wg *sync.WaitGroup) {
				for i := 0; i < numCalls; i++ {
					counterLock.Lock()
					counter++
					counterLock.Unlock()
				}
				wg.Done()
			}(&wg)
		}
		wg.Wait()
	})

	b.Run("atomic", func(b *testing.B) {

		wg.Add(b.N)
		for i := 0; i < b.N; i++ {
			go func(wg *sync.WaitGroup) {
				for i := 0; i < numCalls; i++ {
					atomicCounter.Add(1)
				}
				wg.Done()
			}(&wg)
		}
		wg.Wait()
	})
}

func BenchmarkCounterExplicitInit(b *testing.B) {
	counterLock := sync.Mutex{}
	counter := 0
	atomicCounter := atomic.Int64{}

	b.Run("mutex", func(b *testing.B) {
		wg.Add(b.N)
		for i := 0; i < b.N; i++ {
			go func(wg *sync.WaitGroup) {
				for i := 0; i < numCalls; i++ {
					counterLock.Lock()
					counter++
					counterLock.Unlock()
				}
				wg.Done()
			}(&wg)
		}
		wg.Wait()
	})

	b.Run("atomic", func(b *testing.B) {

		wg.Add(b.N)
		for i := 0; i < b.N; i++ {
			go func(wg *sync.WaitGroup) {
				for i := 0; i < numCalls; i++ {
					atomicCounter.Add(1)
				}
				wg.Done()
			}(&wg)
		}
		wg.Wait()
	})
}

func BenchmarkCounterNoGoroutine(b *testing.B) {
	var counterLock sync.Mutex
	var counter int
	var atomicCounter atomic.Int64

	b.Run("mutex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for i := 0; i < numCalls; i++ {
				counterLock.Lock()
				counter++
				counterLock.Unlock()
			}
		}
	})

	b.Run("atomic", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			for i := 0; i < numCalls; i++ {
				atomicCounter.Add(1)
			}
		}
	})
}
func BenchmarkGlobalCounter(b *testing.B) {

	b.Run("mutex", func(b *testing.B) {
		wg.Add(b.N)
		for i := 0; i < b.N; i++ {
			go func(wg *sync.WaitGroup) {
				for i := 0; i < numCalls; i++ {
					globalCounterLock.Lock()
					globalCounter++
					globalCounterLock.Unlock()
				}
				wg.Done()
			}(&wg)
		}
		wg.Wait()
	})

	b.Run("atomic", func(b *testing.B) {
		wg.Add(b.N)
		for i := 0; i < b.N; i++ {
			go func(wg *sync.WaitGroup) {
				for i := 0; i < numCalls; i++ {
					globalAtomicCounter.Add(1)
				}
				wg.Done()
			}(&wg)
		}
		wg.Wait()
	})
}
