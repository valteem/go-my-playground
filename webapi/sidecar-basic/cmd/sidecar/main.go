package main

import (
	"io"
	"net/http"
)

func proxyHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://127.0.0.1:8080" + r.URL.Path)
	if err != nil {
		http.Error(w, "sidecar failed to get response from main server", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("x-sidecar-response", "sidecar-response")
	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}

func main() {

	http.HandleFunc("/", proxyHandler)
	http.ListenAndServe(":8081", nil)

}
