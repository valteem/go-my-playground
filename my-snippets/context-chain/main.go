package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func AsyncLevelOne(ctx context.Context, wg *sync.WaitGroup) {
//	SyncLevelTwo(ctx, wg)// This leaves no chance for level One default option to run
    go AsyncLevelTwo(ctx, wg)
	for {
		select {
		case <-ctx.Done():
		  log.Println("Sync Level One: context canceled")
		  wg.Done()
		  return
		default:
		  log.Println("Sync Level One: nothing")
		}	
	}
 }

func AsyncLevelTwo(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
		  log.Println("Sync Level Two: context canceled")
		  wg.Done()
		  return
		default:
		  log.Println("Sync Level Two: nothing")
		}
	}
}

func AsyncChain() {
	var wg sync.WaitGroup
	wg.Add(2)
	ctx, cancel := context.WithCancel(context.Background())
	go AsyncLevelOne(ctx, &wg)
	time.Sleep(10 * time.Microsecond)
	cancel()
	wg.Wait()
}

func main() {
	AsyncChain()
}