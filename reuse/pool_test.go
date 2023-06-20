package reuse_test

import (
	"fmt"
	"strconv"
	"sync"
	"testing"

	reuse "github.com/valteem/reuse"
)

func TestPool(t *testing.T) {

	pool := sync.Pool{New: reuse.New}

	for i := 0; i < 5; i++ {
		r := reuse.Resource{}
		r.Name = "name" + strconv.Itoa(i)
		pool.Put(&r) // *reuse.Resource (pointer), not reuse.Resource (object)
	}

	for i := 0; i < 6; i++ {
		r := pool.Get().(*reuse.Resource)
		fmt.Println(r.Name)
	}

}