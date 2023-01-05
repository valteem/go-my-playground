package main

import (
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":3000", http.DefaultServeMux)

}