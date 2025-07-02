package reuse_test

import (
	"context"
	"strings"
	"sync"
	"testing"
)

var (
	output string
	mu     sync.Mutex
)

func apples(ctx context.Context) {
	<-ctx.Done()
	mu.Lock()
	defer mu.Unlock()
	output += "apples"
}

func onions(ctx context.Context) {
	<-ctx.Done()
	mu.Lock()
	defer mu.Unlock()
	output += "onions"
}

func TestFuncTypeAssert(t *testing.T) {

	ch := make(chan any, 2)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		ch <- apples
		wg.Done()
	}()

	go func() {
		ch <- onions
		wg.Done()
	}()

	wg.Wait()
	close(ch)

	for v := range ch {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		v.(func(context.Context))(ctx)
	}

	tokens := []string{"apples", "onions"}

	for _, token := range tokens {
		if !strings.Contains(output, token) {
			t.Errorf("expect output to contain %q, get %q", token, output)
		}
	}

}
