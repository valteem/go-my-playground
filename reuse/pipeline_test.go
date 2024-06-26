// https://go.dev/blog/pipelines

package reuse_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/valteem/reuse"
)

const (
	delayUpperLimit = 5
	numReads = 12
)

func TestSquares(t *testing.T) {

	for v := range reuse.CalcSquares(reuse.MakeSendChan(1, 2, 3, 4)) {
		fmt.Printf("%d ", v)
	} // returns 1, 4, 9, 16
	fmt.Printf("\n")

	go func() {
		for v := range reuse.CalcSquares(reuse.MakeSendChan(10, 20, 30, 40)) {
			fmt.Printf("%d ", v)
		} // returns nothing (without time.Sleep()), looks like gouroutine leak, need further investigation (https://stackoverflow.com/a/64734491)
		fmt.Printf("\n")		
	}()

	time.Sleep(1 * time.Second)
}

func TestSquaresWithDelay(t *testing.T) {
	for i := 1; i <= delayUpperLimit; i++ {
		go func(delay int) {
			time.Sleep(time.Duration(delay) * time.Second)
			for v := range reuse.CalcSquares(reuse.MakeSendChan(100, 200, 300, 400)) {
				fmt.Printf("delay %d output %d\n", delay, v)
			}
		}(i)
	}

	time.Sleep(time.Duration(delayUpperLimit) * time.Second) // roughly 1 sec of delay for every channel
}

func TestSquaresWithWG(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1) // make sure that all sends are completed

	go func() {
		defer wg.Done() // after all sends are completed
		for v := range reuse.CalcSquares(reuse.MakeSendChanWithWG(wg, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)) {
			fmt.Printf("output %d\n", v)
			wg.Done()
		}
	}()

	wg.Wait()

}

func TestUntilDone(t *testing.T) {
	done := make(chan struct{}, 2)
	defer close(done) // sends zero value to _all_ receivers (hence no need for counting receivers)
	s := reuse.SendToChanUntilDone(done)
	r := reuse.ReadFromChanUntilDone(s, done)
	for i := 0; i < numReads; i++ {
		fmt.Println(<-r)
	}
	// done <- struct{}{}
	// done <- struct{}{}

}