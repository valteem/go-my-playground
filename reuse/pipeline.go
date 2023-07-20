// https://go.dev/blog/pipelines

package reuse

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