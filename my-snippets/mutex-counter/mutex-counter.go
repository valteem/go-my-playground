package main

import(
	"fmt"
	"sync"
	"time"
)

const (
	nmax = 100
)

type Counter struct{
	mu sync.Mutex
	count int
}

func (c *Counter) inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) read() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {

	c := Counter{count: 0}

	for i := 0; i < nmax; i++ {
		go c.inc()
		fmt.Println(c.read())
	}

	time.Sleep(time.Second)
	fmt.Println(c.read())

}