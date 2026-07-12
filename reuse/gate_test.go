package reuse_test

import (
	"context"
	"sync/atomic"
	"time"

	"testing"

	"github.com/valteem/reuse"
)

const (
	gateCap   = 2
	numWork   = 14
	workDelay = 10 // time.Milliseconds, use dimensionless value to avoid overflow (time in ns)
)

func TestGateDelay(t *testing.T) {
	g := reuse.NewGate(gateCap)
	workCount := atomic.Int32{}
	now := time.Now()

	for range numWork {
		go func() {
			g.Start(context.Background())
			time.Sleep(workDelay * time.Millisecond) // imitate doing some work
			workCount.Add(1)
			g.Done()
		}()
	}

	for {
		if v := workCount.Load(); v == numWork {
			break
		}
	}

	after := time.Since(now).Milliseconds()
	afterExpected := int64((numWork / gateCap) * workDelay)
	if after < afterExpected {
		t.Errorf("total delay: get %d, expect %d milliseconds", after, afterExpected)
	}

}
