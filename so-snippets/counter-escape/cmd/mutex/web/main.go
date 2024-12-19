// http://localhost:3001/debug/pprof

package main

import (
	"log"
	"net/http"
	"sync"

	_ "net/http/pprof"
)

func main() {

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

	go func() {
		err := http.ListenAndServe("localhost:3001", nil)
		if err != nil {
			log.Fatalf("failed to start pprof endpoint")
		}
	}()

	select {} // block

}
