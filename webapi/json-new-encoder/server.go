package main

import (
	"net/http"
)

func Run(p *Person) {

	mux := http.NewServeMux()
	mux.Handle("/person", HandlePersonData(p))

	http.ListenAndServe(":3001", mux)

}
