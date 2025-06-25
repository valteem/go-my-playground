package main

import (
	"context"
	"net/http"
	"runtime"
	"runtime/pprof"
	"time"

	_ "net/http/pprof"
)

const (
	topLimit = 10_000
)

func main() {

	runtime.SetCPUProfileRate(1_000_000) // just to make it work for this toy example

	chReloadSimple, chShutdownSimple, countSimple := make(chan struct{}), make(chan struct{}), 0
	chReloadTicker, chShutdownTicker, countTicker := make(chan struct{}), make(chan struct{}), 0

	ticker := time.NewTicker(1 * time.Second)

	pprof.Do(context.Background(), pprof.Labels("signal", "ticker"), func(_ context.Context) {
		go func() {
			for {
				select {
				case <-chShutdownTicker:
					return
				case <-ticker.C:
					select {
					case <-chReloadTicker:
						reload(&countTicker)
					case <-chShutdownTicker:
						return
					}
				}
			}
		}()
	})

	pprof.Do(context.Background(), pprof.Labels("signal", "simple"), func(_ context.Context) {
		go func() {
			for {
				select {
				case <-chReloadSimple:
					reload(&countSimple)
				case <-chShutdownSimple:
					return
				}
			}
		}()
	})

	go func() {
		for {
			time.Sleep(50 * time.Millisecond)
			chReloadSimple <- struct{}{}
			chReloadTicker <- struct{}{}
		}
	}()

	http.ListenAndServe(":6060", nil)

}

func reload(count *int) {

	for {
		if *count == topLimit {
			*count = 0
			return
		}
		*count++
	}

}
