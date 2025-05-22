package main

import (
	"net/http"
)

func main() {

	// fs := http.FileServer(http.Dir("static"))

	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	// http.ServeMux.Handle() and http.Dir() and sum up their paths
	// https://stackoverflow.com/a/74969884
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":3001", nil)
}
