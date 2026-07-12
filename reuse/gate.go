package reuse

import (
	"context"
)

type Gate struct {
	ch chan struct{}
}

func NewGate(cap int) *Gate {
	return &Gate{ch: make(chan struct{}, cap)}
}

func (g *Gate) Start(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case g.ch <- struct{}{}:
		// if channel is full - waiting until some reads happen at the other end via Done()
		return nil
	}
}

func (g *Gate) Done() {
	select {
	case <-g.ch:
	default:
		panic("gate.Done: more ops done than started")
	}
}
