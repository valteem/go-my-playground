package main

import (
	"fmt"
	"sync"
	"unsafe"
)

func worker(nw int, nj int, wg *sync.WaitGroup) {

	for count := 1; count <= nj; count ++ {
        fmt.Printf("Worker #%d performing job #%d\n", nw, count)
	}
	wg.Done()
}

func main() {
	
	numWorkers := 15
	numJobs := 2

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for workerCount := 1; workerCount <= numWorkers; workerCount++ {
		go worker(workerCount, numJobs, &wg)
	}

	wg.Wait()

}