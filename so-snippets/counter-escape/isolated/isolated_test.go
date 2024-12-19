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
