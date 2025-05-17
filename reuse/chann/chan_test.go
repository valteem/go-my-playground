package chann

import (
	"sync"
	"testing"
)

func TestReadOK(t *testing.T) {

	numWrites := 5
	wg := sync.WaitGroup{}
	wg.Add(numWrites)

	ch := make(chan int)
	go func() {
		for i := range numWrites {
			ch <- i
			wg.Done()
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	// drain channel in case of premature return
	defer func() {
		if ch != nil {
			for range ch {
			}
		}
	}()

	chcopy := ch
	output := []int{}
	phantomReads := 0

	for {
		select {
		case v, ok := <-chcopy:
			if !ok {
				chcopy = nil
				break
			}
			output = append(output, v)
		default:
			phantomReads++ // IRL should be replaced with adding more workers to process chcopy output
		}
		if chcopy == nil { // reason to use chcopy instead of ch
			break
		}

	}

	if len(output) != numWrites {
		t.Errorf("output: get %d values, expect %d", len(output), numWrites)
	}

}

func TestReadOKWithFlag(t *testing.T) {

	numWrites := 5
	wg := sync.WaitGroup{}
	wg.Add(numWrites)

	ch := make(chan int)
	go func() {
		for i := range numWrites {
			ch <- i
			wg.Done()
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	// drain channel in case of premature return
	defer func() {
		if ch != nil {
			for range ch {
			}
		}
	}()

	var flag bool
	output := []int{}
	phantomReads := 0

	for {
		select {
		case v, ok := <-ch:
			if !ok {
				// replace nil out of copied channel reference to bool flag
				flag = true
				break
			}
			output = append(output, v)
		default:
			phantomReads++
		}
		if flag {
			break
		}

	}

	if len(output) != numWrites {
		t.Errorf("output: get %d values, expect %d", len(output), numWrites)
	}

}
