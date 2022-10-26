package main

import "net/http"

func main() {

	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Some plain text")) // looks like Content Type text/plain is by default

}