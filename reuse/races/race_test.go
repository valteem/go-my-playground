// go test -races github.com/valteem/reuse/races

package races

import (
	"sync"
	"testing"
)

const (
	numWorkers                 = 1000
	numCallsToIncrementCounter = 1000
)

var (
	count = 0
)

func TestRaceDetection(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(wg *sync.WaitGroup) {
			for j := 0; j < numCallsToIncrementCounter; j++ {
				count++
			}
			wg.Done()
		}(&wg)
	}

	wg.Wait()

	if count == numWorkers*numCallsToIncrementCounter {
		t.Errorf("expect race condition, found nothing")
	}

}
