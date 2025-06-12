package initorder

import (
	"context"
	"fmt"
	"sync"
)

func init() {
	fmt.Println("init server")
}

func runServer(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("running server")
	<-ctx.Done()
	wg.Done()
	fmt.Println("quitting server")
}
