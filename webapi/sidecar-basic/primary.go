package main

import (
	"io"
	"net/http"
	"strings"
)

func primaryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Primary-Status", "primary-running")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, strings.NewReader("primary service response"))
}

func runPrimary() {

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(primaryHandler))
	http.ListenAndServe(":3001", mux)

}
