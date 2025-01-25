package reuse_test

import (
	"testing"
	"time"
)

func TestLoopCond(t *testing.T) {

	cond := false
	now := time.Now()
	var since time.Duration

	go func() {
		for !cond {
		}
		since = time.Since(now)
	}()

	time.Sleep(1 * time.Second)
	cond = true

	// Allow some more time before testing since value
	time.Sleep(100 * time.Millisecond)

	if actual, expected := since, 1*time.Second; actual < expected {
		t.Errorf("time.Since(): get %d, expect >= %d", actual, expected)
	}

}
