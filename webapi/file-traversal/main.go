// https://stackoverflow.com/a/58655915 and comment that follows

package main

import (
	"net/http"
)

func serveFile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		path = "./assets/index.html"
	}
	http.ServeFile(w, r, path)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(serveFile))
	http.ListenAndServe(":8080", mux)
}

// 127.0.0.1:8080/opt/google/chrome returns content of '/opt/google/chrome' directory

// https://www.stackhawk.com/blog/golang-path-traversal-guide-examples-and-prevention/
