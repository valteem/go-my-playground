package initorder

import (
	"context"
	"fmt"
	"sync"
)

func init() {
	fmt.Println("init client")
}

func runClient(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("running client")
	<-ctx.Done()
	wg.Done()
	fmt.Println("quitting client")
}
