// Implement ring buffer with generics
// Mostly follows smallnest/ringbuffer mechanics

package ringbuf

import (
	"sync"
)

type RingBuf[T any] struct {
	buf []T

	readCond sync.Cond // https://victoriametrics.com/blog/go-sync-cond/
}
