// https://go.dev/blog/pipelines

package reuse

import (
//	"fmt"
	"sync"
)

func WriteChan(ch chan int, s []int) {
	for _, v := range s {
//		fmt.Println("input:", v)
		ch <- v
	}
	close(ch)
}

func MergeChannels(chs ...<-chan int) <-chan int {

	mch := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(len(chs))
		for _, c := range chs {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for v := range ch {
//					fmt.Println("read:", v)
					mch <- v
				}	
			}(c, wg)
		}
	go func() {
		wg.Wait()
		close(mch)	
	}()
	}()

	return mch

}