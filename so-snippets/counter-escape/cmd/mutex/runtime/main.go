/*
go build
./runtime
go tool pprof runtime mem.prof
list main.main

Memory profile is empty without starting CPU profiling in advance
*/

package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

func main() {

	cpuProfile, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(cpuProfile)
	defer pprof.StopCPUProfile()

	lock := sync.Mutex{}
	counter := 0

	go func() {
		for {
			lock.Lock()
			counter++
			lock.Unlock()
		}
	}()

	go func() {
		for {
			lock.Lock()
			counter--
			lock.Unlock()
		}
	}()

	runtime.GC()
	memProfile, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer memProfile.Close()
	if err := pprof.WriteHeapProfile(memProfile); err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)

}
