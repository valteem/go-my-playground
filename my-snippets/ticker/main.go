package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

const (
	topLimit               = 1_000_000
	uptimeMilliSec         = 1000
	tickerIntervalMilliSec = 200
	reloadIntervalMilliSec = 50
	profileRateHz          = 10_000
)

func main() {

	f, err := os.Create("cpu.prof.pb.gz")
	if err != nil {
		log.Fatalf("failed to create profile file")
	}
	defer f.Close()

	runtime.SetCPUProfileRate(profileRateHz) // standard 100 Hz is not enough

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatalf("failed to start profiling: %v", err)
	}
	defer pprof.StopCPUProfile()

	payloadSimple, payloadTicker := 0, 0
	countReloadSimple, countReloadTicker := 0, 0

	ticker := time.NewTicker(tickerIntervalMilliSec * time.Millisecond)
	tickerReloadWithTicker := time.NewTicker(reloadIntervalMilliSec * time.Millisecond)
	tickerReloadSimple := time.NewTicker(reloadIntervalMilliSec * time.Millisecond)
	tickerUptimeWithTicker := time.NewTicker(uptimeMilliSec * time.Millisecond)
	tickerUptimeSimple := time.NewTicker(uptimeMilliSec * time.Millisecond)

	var wg sync.WaitGroup
	wg.Add(2)

	pprof.Do(context.Background(), pprof.Labels("signal", "with ticker"), func(_ context.Context) {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-tickerUptimeWithTicker.C:
					return
				case <-ticker.C:
					select {
					case <-tickerReloadWithTicker.C:
						reload(&payloadTicker)
						countReloadTicker++
					case <-tickerUptimeWithTicker.C:
						return
					}
				}
			}
		}()
	})

	pprof.Do(context.Background(), pprof.Labels("signal", "simple"), func(_ context.Context) {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-tickerReloadSimple.C:
					reload(&payloadSimple)
					countReloadSimple++
				case <-tickerUptimeSimple.C:
					return
				}
			}
		}()
	})

	wg.Wait()

	fmt.Printf("ticker reloaded %d times, simple reloaded %d times\n", countReloadTicker, countReloadSimple)

}

func reload(payload *int) {

	for {
		if *payload == topLimit {
			*payload = 0
			return
		}
		*payload++
	}

}
