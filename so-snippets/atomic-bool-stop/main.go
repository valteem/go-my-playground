package main

import (
    "sync/atomic"
    "time"

	"myexample/internal"
)

func main() {
    var stop atomic.Bool
    go internal.AsyncForeverTask(&stop)
    time.Sleep(10 * time.Second)
    stop.Store(true)
}