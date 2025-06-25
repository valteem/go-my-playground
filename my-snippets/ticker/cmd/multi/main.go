package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"sync"
)

func averageRand(n int, output *float64) {
	*output = 0.
	var kfloat float64
	for k := 1; k <= n; k++ {
		kfloat = float64(k)
		*output = ((kfloat-1.)*(*output) + rand.Float64()) / kfloat
	}
}

func main() {

	f, err := os.Create("cpu.prof.pb.gz")
	if err != nil {
		log.Fatalf("failed to create profile file: %v", err)
	}
	defer f.Close()

	runtime.SetCPUProfileRate(10_000) // just to make it work for this toy example
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatalf("failed to start profiling: %v", err)
	}
	defer pprof.StopCPUProfile()

	sizes := []int{1_000_000, 2_000_000, 3_000_000, 4_000_000}

	var wg sync.WaitGroup
	wg.Add(len(sizes))

	for _, r := range sizes {
		pprof.Do(context.Background(), pprof.Labels("sample size", strconv.Itoa(r)), func(_ context.Context) {
			var output float64
			go func() {
				averageRand(r, &output)
				wg.Done()
			}()
		})
	}

	wg.Wait()

}
