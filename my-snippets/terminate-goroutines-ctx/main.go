package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Work struct{
	attr string // just for illustration, not used anywhere 
}

func (w *Work) Check(ctx context.Context, workDuration time.Duration) bool {
	select {
// ctx.Done() returns a channel that receives an empty struct{} type every time the context receives cancellation event		
		case <- ctx.Done(): 
			return false
// time.After() return a channel that waits for duration to elapse and then receives the current time
		case <- time.After(workDuration):
			return true
		case <- time.After(3000 * time.Millisecond):
			return false
	} 
}

func CheckAll(works []*Work) {

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // probably just in case cancel() is not invoked anywhere later

	num := len(works)
	results := make(chan bool, num)

	for i, w := range works {
		workDuration := time.Second * time.Duration(i) // type conversion (to nanoseconds)
		wg.Add(1)
		go func(w *Work) {
			defer wg.Done()
			fmt.Println(workDuration.Seconds())
			result := w.Check(ctx, workDuration)
			select {
			case results <- result:
			case <-ctx.Done():
				return
			}
		}(w)
	}

/* 	go func() {
		wg.Wait()
		close(results)
	}() */

	for result := range results {
		fmt.Println("Result", result)
		if !result {
			cancel()
			break
		}
	}

// This works too, without wrapping it in a goroutine	
	wg.Wait()
	close(results)
}

func main() {
	CheckAll(make([]*Work, 10))
}