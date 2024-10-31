package main

import (
	"io"
	"net/http"
	"strings"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("x-main-header", "main-response")
	io.Copy(w, strings.NewReader("main handler response\n"))
}

func main() {

	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)

}
