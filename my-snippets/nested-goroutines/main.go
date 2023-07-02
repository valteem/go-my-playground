package main

import (
	"fmt"
	"sync"
)

func send(c chan int, nmax int, i int) {
	i++
	if i > nmax {
		return
	} else {
		fmt.Println("sending ...", i)
		c <- i
		go send(c, nmax, i)
		return
	}
}

const (
	nMax = 100
)

func main() {
	c := make(chan int, nMax)
	var wg sync.WaitGroup
	wg.Add(nMax)
	go send(c, nMax, 0)
	readCount := 0;
	go func(){
		for {
			select {
			case i := <- c:
				fmt.Println("reading ...", i)
				readCount++
				wg.Done()
				if i == nMax {return}
			default:
				fmt.Println("nothing ...")
			}	
		}
	}()
	wg.Wait()
	fmt.Println("read count", readCount)
}