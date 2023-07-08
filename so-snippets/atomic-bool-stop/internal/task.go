package internal

import (
    "fmt"
    "sync/atomic"
    "time"
)

func AsyncForeverTask(stop *atomic.Bool) {
    for {
        if stop.Load() {
            return
        }
        fmt.Println("I'm still running")
        time.Sleep(1 * time.Second)
    }
}
