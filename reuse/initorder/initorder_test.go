package initorder

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestInitOrder(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	runServer(ctx, wg)
	runClient(ctx, wg)

	cancel()

	wg.Wait()

}
