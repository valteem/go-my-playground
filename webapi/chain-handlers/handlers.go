package main

import (
	"log"
	"net/http"
)

func withAddHeaderA(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Custom-Header-A", "custom-header-a-value")
		h.ServeHTTP(w, r)
	})
}

func withAddHeaderB(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Custom-Header-B", "custom-header-b-value")
		h.ServeHTTP(w, r)
	})
}

func withAddHeaderC(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Custom-Header-C", "custom-header-c-value")
		h.ServeHTTP(w, r)
	})
}

func runServer(h http.Handler) {

	mux := http.NewServeMux()

	mux.Handle("GET /headers", withAddHeaderA(withAddHeaderB(withAddHeaderC(h))))

	err := http.ListenAndServe(":3001", mux)
	if err != nil {
		log.Fatalf("failed to listen and serve: %v", err)
	}

}
