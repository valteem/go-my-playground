package main

import (
	"fmt"
	"sync"
)

type WriterToChan struct {
	ch chan string
}

func (wtc WriterToChan) Write(p []byte) (int, error) {
	wtc.ch <- string(p)
	return len(p), nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan string)
	q := make(chan bool)
	go func() {
		for {
			select{
			case <- q:
				wg.Done()
				return
			default:
				fmt.Println("This is goroutine, input through channel is", <- ch) // this one reads from ch
			}
		}
	}()
	wtc := WriterToChan{ch:ch}
	wtc.Write([]byte("some text")) // this eventually writes to ch
	fmt.Fprintf(wtc, "%s", "another text") // this one too
	q <- true
	wg.Wait()
}