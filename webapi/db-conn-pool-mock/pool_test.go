package main

import (
	"io"
	//	"log"
	"net/http"
	"slices"
	"strconv"
	"sync"
	"testing"
)

const (
	reqNum = 10000
)

func runServer() {

	mux := http.NewServeMux()
	db := &DB{count: 0, mu: &sync.Mutex{}}
	h := NewHandler(db)

	mux.Handle("/count", http.HandlerFunc(h.GetCount))

	http.ListenAndServe(":3001", mux)

}
func TestPool(t *testing.T) {

	counts := make(chan int, reqNum)
	var expected, output []int

	go runServer()

	for i := 0; i < reqNum; i++ {
		expected = append(expected, i+1)
		go func() {
			resp, err := http.Get("http://localhost:3001/count")
			// TODO: add error processing within 'non-test' goroutines
			if err != nil {
				//			log.Println(err)
			}
			defer resp.Body.Close()
			b, _ := io.ReadAll(resp.Body)
			count, _ := strconv.Atoi(string(b))
			counts <- count
		}()
	}

	var wg sync.WaitGroup
	wg.Add(reqNum)

	go func() {
		for count := range counts {
			output = append(output, count)
			wg.Done()
		}
	}()

	wg.Wait()

	slices.Sort(output)

	if !slices.Equal(output, expected) {
		t.Errorf("expect sorted int slice")
	}

}
