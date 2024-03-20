package timeutil

import (
	"sync"
	"time"
)

type Clock interface {
	Now() time.Time
}

type realTimeClock struct{}

func (_ *realTimeClock) Now() time.Time {
	return time.Now()
}

func NewRealClock() Clock {
	return &realTimeClock{} // returning pointer cause there is no `func (_ realTimeClock) Now()` (only `(_ *realTimeClock)`)
}

type SimulatedClock struct {
	mu sync.Mutex
	t  time.Time
}

func (c *SimulatedClock) Now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.t
}

func (c *SimulatedClock) SetTime(t time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.t = t
}

func (c *SimulatedClock) AdvanceTime(d time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.t = c.t.Add(d)
}

func NewSimulatedClock(t time.Time) *SimulatedClock {
	return &SimulatedClock{t: t}
}
