package main

import (
    "fmt"
    "time"
)

var done bool = false

func asyncForeverTask() {
    for {
        if done {
            return
        }
        fmt.Println("I'm still running")
        time.Sleep(1 * time.Second)
    }
}

func main() {
    go asyncForeverTask()
    time.Sleep(10 * time.Second)
    done = true
}