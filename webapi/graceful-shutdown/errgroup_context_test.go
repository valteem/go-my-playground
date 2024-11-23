// https://gist.github.com/pteich/c0bb58b0b7c8af7cc6a689dd0d3d26ef
package main

import (
	"context"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	egNum        = 100
	spawnTimeout = 5 * time.Second
	trunkTimeout = 1 * time.Second
)

func TestErrgroupCtx(t *testing.T) {

	chAfter, chDone := make(chan int, egNum), make(chan int, egNum)

	ctx, cancel := context.WithCancel(context.Background())

	eg, ctxEg := errgroup.WithContext(ctx)

	i := 0
	for i < egNum {
		eg.Go(func() error {
			select {
			case <-time.After(spawnTimeout):
				chAfter <- 1
				return nil
			case <-ctxEg.Done():
				chDone <- 1
				return ctxEg.Err()
			}
		})
		i++
	}

	time.Sleep(trunkTimeout)

	cancel()

	eg.Wait()
	close(chAfter)
	close(chDone)

	var countAfter, countDone int
	for range chAfter {
		countAfter++
	}
	for range chDone {
		countDone++
	}

	if countDone != egNum {
		t.Errorf("goroutines closed by <-ctx.Done / by <-time.After():\nget %d / %d, expect %d / 0", countDone, countAfter, egNum)
	}

}
