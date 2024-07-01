// go test -memprofile mem.prof -bench examine -run BenchmarkPool
// go tool pprof mem.prof
package examine

import (
	"fmt"
	"sync"
	"testing"
)

type heavyObject struct {
	s []string
}

func NewHeavyObject(cap int, str string) *heavyObject {
	h := &heavyObject{s: make([]string, 0, cap)}
	for i := 0; i < cap; i++ {
		h.s = append(h.s, str)
	}
	return h
}
func BenchmarkPool(b *testing.B) {

	benchmarks := []struct {
		str  string
		size int
		cap  int
	}{
		{
			str:  "a`",
			size: 1000 * 1000,
			cap:  100 * 10,
		},
		{
			str:  "a",
			size: 1000 * 1000 * 10,
			cap:  100,
		},
	}

	for _, bm := range benchmarks {
		benchmarkName := fmt.Sprintf("sync.pool_size_%d_capacity_%d", bm.size, bm.cap)
		b.Run(benchmarkName, func(b *testing.B) {
			p := &sync.Pool{New: func() any {
				return NewHeavyObject(bm.cap, bm.str)
			},
			}
			ch := make(chan *heavyObject, bm.size)
			wg := sync.WaitGroup{}
			wg.Add(bm.size)

			go func() {
				for i := 0; i < bm.size; i++ {
					// getting objects from pool
					hObj, _ := p.Get().(*heavyObject)
					ch <- hObj
				}
			}()

			// putting objects back to the pool
			go func() {
				for hObj := range ch {
					p.Put(hObj)
					wg.Done()
				}
			}()
			wg.Wait()
		})
	}
}
