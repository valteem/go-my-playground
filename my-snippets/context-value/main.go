package main

import (
	"context"
	"fmt"
	"sync"
)

type contextKey string

// https://pkg.go.dev/context#pkg-overview
// Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions
// This example seemingly does exactly the opposite
func Waiting(ctx context.Context, key contextKey, wg *sync.WaitGroup) {
	for {
		if ctx.Value(key) == "exit" {
			fmt.Println("exit signal received")
			wg.Done()
			return
		}
	}
}

func main () {

	var wg sync.WaitGroup
	wg.Add(1)
	ctxKey := contextKey("signal")
	ctx := context.WithValue(context.Background(), ctxKey, "exit")
	go Waiting(ctx, ctxKey, &wg)
	wg.Wait()

}