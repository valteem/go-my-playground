package reuse_test

import (
	"sync"
	"testing"
)

func TestChannelRead(t *testing.T) {

	c := make(chan int, 5)

	for i := 0; i < 5; i++ {
		c <- i
	}

	for i := 0; i < 5; i++ {
		j, ok := <-c
		if j != i || !ok {
			t.Errorf("reading from buffered channel: get %d, %t, expect %d, %t", j, ok, i, true)
		}
	}

	// trying to read from empty (but not yet closed) channel blocks goroutine

	close(c)

	for i := 0; i < 5; i++ {
		j, ok := <-c
		if j != 0 || ok {
			t.Errorf("reading from closed channel: get %d, %t, expect %d, %t", j, ok, 0, false)
		}
	}

}

func TestReadFromClosedUnbufferedChannel(t *testing.T) {

	var (
		numSend = 100
		numRead = 500
	)

	sink := []int{}

	var wg sync.WaitGroup
	wg.Add(numSend + numRead)

	ch := make(chan int)

	go func() {
		for i := range numSend {
			ch <- i
			wg.Done()
		}
		close(ch)
	}()

	go func() {
		for i := 0; i < numRead; i++ {
			sink = append(sink, <-ch)
			wg.Done()
		}
	}()

	wg.Wait()

	if len(sink) != numRead {
		t.Errorf("sink size: get %d, expect %d", len(sink), numRead)
	}

}

func TestReadFromClosedBufferedChannel(t *testing.T) {

	var (
		numSend = 100
		numRead = 500
		bufSize = 10
	)

	sink := []int{}

	var wg sync.WaitGroup
	wg.Add(numSend + numRead)

	ch := make(chan int, bufSize)

	go func() {
		for i := range numSend {
			ch <- i
			wg.Done()
		}
		close(ch)
	}()

	go func() {
		for i := 0; i < numRead; i++ {
			sink = append(sink, <-ch)
			wg.Done()
		}
	}()

	wg.Wait()

	if len(sink) != numRead {
		t.Errorf("sink size: get %d, expect %d", len(sink), numRead)
	}

}
