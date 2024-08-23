package routes

import (
	"net/http"
)

func someHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("the path is " + r.URL.Path))
}

func someOtherHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("request path is " + r.URL.Path))
}
