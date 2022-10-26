package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request){
		tm := time.Now().Format(format)
		w.Write([]byte("The time is " + tm))
	}
	return http.HandlerFunc(fn)
}

func main() {

	mux := http.NewServeMux()

	th := timeHandler(time.RFC1123)

	mux.Handle("/time", th) // shortcut for mux.Handle(pattern, HandlerFunc(handler))

	log.Println("Listening ...")

	http.ListenAndServe(":3000", mux)
}