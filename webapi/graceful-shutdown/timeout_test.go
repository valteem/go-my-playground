package main

import (
	"context"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

const (
	numReq = 100
)

func handleTimeout(ctx context.Context, timeout time.Duration, chDone, chAfter chan int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-ctx.Done():
			chDone <- 1
			return
		case <-time.After(timeout):
			chAfter <- 1
			return
		}
	})
}

func TestTimeout(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	timeout := 2 * time.Second
	chDone, chAfter := make(chan int, numReq), make(chan int, numReq)

	mux := http.NewServeMux()
	mux.Handle("/timeout", handleTimeout(ctx, timeout, chDone, chAfter))

	go func() {
		http.ListenAndServe(":3001", mux)
	}()

	time.Sleep(1 * time.Second) // allow server some time to start

	var wg sync.WaitGroup
	wg.Add(numReq)

	for i := 0; i < numReq; i++ {
		go func(i int) {
			req, err := http.NewRequest(http.MethodGet, "http://localhost:3001/timeout", nil)
			if err != nil {
				log.Printf("failed to create request #%d", i)
			}
			client := http.Client{}
			_, err = client.Do(req)
			if err != nil {
				log.Printf("failed to get response #%d: %v", i, err)
			}
		}(i)
		wg.Done()
	}

	wg.Wait()
	cancel()

	time.Sleep(100 * time.Millisecond) // allow some time for cancel() to propagate

	close(chDone)
	close(chAfter)

	var countDone, countAfter []int
	for count := range chDone {
		countDone = append(countDone, count)
	}
	for count := range chAfter {
		countAfter = append(countAfter, count)
	}

	if len(countDone) < numReq {
		t.Errorf("expect all %d requests completed by ctx.Done(), get %d", numReq, len(countDone))
	}

	if len(countAfter) != 0 {
		t.Errorf("expect no requests completed by time.After(), get %d", len(countAfter))
	}

}
