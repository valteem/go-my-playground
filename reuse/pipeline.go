// https://go.dev/blog/pipelines

package reuse

import "sync"

func MakeSendChan(num ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range num {
			out <- v
		}
		close(out)
	}()
	return out
}

func CalcSquares(ch <-chan int) <-chan int {
	squares := make(chan int)
	go func() {
		for v := range ch {
			squares <- v * v
		}
		close(squares)
	}()
	return squares
}

func MakeSendChanWithWG(wg *sync.WaitGroup, num ...int) <- chan int {
	out := make(chan int)
	wg.Add(len(num)) // make sure wg counter is set to max value before receiver in main/test starts reading from channel
	go func() {
		for _, v := range num {
			out <- v
//			wg.Add(1) // this way wg counter in main/test may reach zero before all sends are completed
		}
		close(out)
	}()
	return out	
}

func SendToChanUntilDone(done <-chan struct{}) <-chan int {
	out := make(chan int)
	i := 1
	go func() {
		for {
			select {
			case <-done:
				close(out)
				return
			default:
				out <- i
				i++
			}
		}
	}()
	return out
}

func ReadFromChanUntilDone(ch <-chan int, done chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		for {
			select {
			case <-done:
				close(out)
				return
			case v := <-ch:
				out <- v
			}
		}
	}()
	return(out)
}