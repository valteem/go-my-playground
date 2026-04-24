package reuse

import (
	"sync"
	"time"

	"testing"
)

const (
	resourceSleepTime = time.Duration(100 * time.Nanosecond)
)

type localResourceGuard struct {
	s  bool
	mu sync.Mutex
}

func TestWaitingOnResource(t *testing.T) {

	g := localResourceGuard{s: false}

	now := time.Now()

	go func() {
		g.mu.Lock()
		time.Sleep(resourceSleepTime)
		g.s = true
		g.mu.Unlock()
	}()

	for {
		// https://github.com/openeverest/openeverest-operator/blob/f18498196ba135936ef677e49ee478aca9c79390/internal/controller/everest/common/helper.go#L826
		// looks like quite clumsy drop in replacement for WaitGroup
		if !g.s {
			continue
		} else {
			break
		}
	}

	delay := time.Since(now)
	if delay < resourceSleepTime {
		t.Errorf("waiting on resource: get %s, expect no less than %s", delay, resourceSleepTime)
	}

}
