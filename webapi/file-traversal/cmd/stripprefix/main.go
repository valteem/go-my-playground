package main

import (
	"net/http"
)

func main() {

	http.Handle("/assets", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":3001", nil)
}
