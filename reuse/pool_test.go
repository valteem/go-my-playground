package reuse_test

import (
	"fmt"
	"sync"
	"testing"

	reuse "github.com/valteem/reuse"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{New: reuse.New}
	testResource := pool.Get().(*reuse.Resource)
	fmt.Println(testResource.Name)
	testResource.Name = "another resource"
	pool.Put(testResource)
	anotherResource :=  pool.Get().(*reuse.Resource)
	fmt.Println(anotherResource.Name)
	pool.Put(anotherResource)
	yetAnotherResource :=  pool.Get().(*reuse.Resource)
	fmt.Println(yetAnotherResource.Name)
	yetAnotherResource.Name = "yet another resource"
	pool.Put(yetAnotherResource)
	for i := 0; i < 4; i++ {
		more :=  pool.Get().(*reuse.Resource)
		fmt.Println(more.Name)
	}
}