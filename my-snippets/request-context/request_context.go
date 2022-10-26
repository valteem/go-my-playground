package main

import (
    "fmt"
    "net/http"
    "time"
)

func hello(w http.ResponseWriter, req *http.Request) {

    ctx := req.Context()
    fmt.Println("server: hello handler started")
    defer fmt.Println("server: hello handler ended")

    for {
        select {
        case <-time.After(2 * time.Second):
            fmt.Fprintf(w, "hello\n")
            return
        case <-ctx.Done(): // no idea how to change request context, so this entry is unreachable
            err := ctx.Err()
            fmt.Println("server:", err)
            internalError := http.StatusInternalServerError
            http.Error(w, err.Error(), internalError)
            return
        }
    }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":8090", nil)

}