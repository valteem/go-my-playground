package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is " + tm))
}

func main() {

	mux := http.NewServeMux()

	th := http.HandlerFunc(timeHandler) // looks like a decorator

	mux.Handle("/time", th)

	log.Println("Listening ...")

	http.ListenAndServe(":3000", mux)
}