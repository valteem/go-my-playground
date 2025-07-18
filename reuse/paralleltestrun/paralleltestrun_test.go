package paralleltestrun

import (
	"fmt"
	"testing"
	"time"
)

func TestNestedRuns(t *testing.T) {

	delayUnit := 100 * time.Millisecond

	for i := range 5 {
		delay := delayUnit * time.Duration(5-i)
		now := time.Now()
		t.Run(fmt.Sprintf("test run #%d with output delay of %v", i, delay), func(t *testing.T) {
			t.Parallel() // makes test actually run in parallel
			time.Sleep(delay)
			fmt.Println(i, delay, time.Since(now))
		})
	}

}
