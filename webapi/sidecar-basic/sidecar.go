package main

import (
	"io"
	"net/http"
)

func proxyHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://127.0.0.1:3001" + r.URL.Path)
	if err != nil {
		http.Error(w, "sidecar failed to get response from primary server", http.StatusInternalServerError)
		return
	}
	w.Header().Set("X-Sidecar-Status", "sidecar-running")
	primaryHeader, ok := resp.Header["X-Primary-Status"]
	if ok {
		w.Header().Set("X-Primary-Status", primaryHeader[0])
	}
	defer resp.Body.Close()
	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}

func runProxy() {

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(proxyHandler))
	http.ListenAndServe(":3002", mux)

}
