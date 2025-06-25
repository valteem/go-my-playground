package main

import (
	"context"
	"net/http"
	"runtime/pprof"

	_ "net/http/pprof"
)

const (
	topLimit = 10_000
)

func main() {

	pprof.Do(context.Background(), pprof.Labels("count", "1st"), func(_ context.Context) {
		go func() {
			count := 0
			for {
				reload(&count)
			}
		}()
	})

	pprof.Do(context.Background(), pprof.Labels("count", "2nd"), func(_ context.Context) {
		go func() {
			count := 0
			for {
				reload(&count)
			}

		}()
	})

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
