package reuse_test

import (
	"sync"
	"testing"
)

func TestChanSync(t *testing.T) {

	var counter int

	const level int = 10000

	chSync := make(chan struct{})

	go func() {
		for counter < level {
			counter++
		}
		close(chSync)
	}()

	<-chSync

	if counter != level {
		t.Errorf("counter value: get %d, expect %d", counter, level)
	}

}

func TestChanMultiSync(t *testing.T) {

	var counter int
	var mtx sync.RWMutex
	const nWorkers int = 10
	const maxCounterValue = 100000
	chSync := make(chan struct{}, nWorkers)

	for i := 0; i < nWorkers; i++ {
		go func(i int) {
			for {
				mtx.RLock()
				if counter < maxCounterValue {
					counter++
					mtx.RUnlock()
				} else {
					mtx.RUnlock()
					break
				}
			}
			chSync <- struct{}{}
		}(i)
	}

	for len(chSync) < nWorkers {
	}

	if counter != maxCounterValue {
		t.Errorf("counter value: get %d, expect %d", counter, maxCounterValue)
	}

}
