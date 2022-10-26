package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container {
		counters: map[string]int{"a": 0, "b":0}, // эта запятая нужна потому, что иначе компилятор поставит сюда точку с запятой
		                                         // и получатся незакрытые скобки
												 // https://stackoverflow.com/a/62576035
	}

	var wg sync.WaitGroup

	doInc := func (name string, num int) {
		for i := 0; i < num; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(4)
	go doInc("a", 10000)
	go doInc("a", 5000)
	go doInc("b", 30000)
	go doInc("b", 3000)

	wg.Wait()

	fmt.Println(c.counters)

}