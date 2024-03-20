package timeutil

import (
	"testing"
	"time"
)

func TestSimulatedClock(t *testing.T) {
	now := time.Now()
	tests := []struct {
		description  string
		startTime    time.Time
		timeStep     time.Duration
		expectedTime time.Time
	}{
		{
			description:  "time step forward",
			startTime:    now,
			timeStep:     60 * time.Second,
			expectedTime: now.Add(60 * time.Second),
		},
		{
			description:  "time step backward",
			startTime:    now,
			timeStep:     -60 * time.Second,
			expectedTime: now.Add(-60 * time.Second),
		},
	}
	for _, tst := range tests {
		c := NewSimulatedClock(tst.startTime)
		if c.Now() != tst.startTime {
			t.Errorf("%s: Before time step; SimulatedClock.Now() = %+v, expect %+v", tst.description, c.Now(), tst.startTime)
		}
	}
}
