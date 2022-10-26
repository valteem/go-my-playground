package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	ctx := context.Background()
	ctx, cancelCtx := context.WithCancel(ctx)
	defer fmt.Println("Finished")

	go func() {
		defer wg.Done()
		count := 0
		for {
			select {
			case <-time.After(1 * time.Millisecond):
				fmt.Printf("Emitting cancellation signal after %d counts\n", count)
				cancelCtx()
				return
			default:
			 	count++
			}
		}
	}()

	go func() {
		defer wg.Done()
		count :=- 0
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Cancellation signal detected, quitting after %d counts\n", count)
				return
			default:
				count++
			}
		}
	}()

	wg.Wait()

}