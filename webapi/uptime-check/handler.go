package main

import (
	"net/http"
)

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("service is up and running"))
	w.WriteHeader(http.StatusOK)
}

func runHandler() {

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(basicHandler))

	http.ListenAndServe(":3001", mux)
}
