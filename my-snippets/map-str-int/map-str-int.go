package main

import (
	"fmt"
	"strconv"
)

type counter struct{
	v map[string]int
}

const (
	maxnum = 5
)

func main() {

	var ckey string

	c := counter{v : make(map[string]int)}
	for i := 0; i < maxnum; i++ {
		ckey = "key" + strconv.Itoa(i)
		c.v[ckey] = i
		fmt.Println(c)
	}
}