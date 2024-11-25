package handlers

import (
	"net/http"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Wrap-Auth", "hello")
}

func HandleInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Wrap-Auth", "info")
}
