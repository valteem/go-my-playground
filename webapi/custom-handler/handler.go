package main

import (
	"io"
	"net/http"
)

type CustomHandler struct {
	Msg          map[string]string
	BasicHandler http.ServeMux
}

func New() *CustomHandler {
	return &CustomHandler{Msg: make(map[string]string), BasicHandler: *http.NewServeMux()}
}

func (ch *CustomHandler) AddMessage(key, value string) {
	ch.Msg[key] = value
}

func (ch *CustomHandler) Handle(pattern string, handler http.Handler) {
	ch.BasicHandler.Handle(pattern, handler)
}

func (ch *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ch.BasicHandler.ServeHTTP(w, r)
}

// http.ServeMux contains lock, thus need for pointer receiver
func (ch *CustomHandler) handleMsg() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.PathValue("keyID")
		// It does not even get there if keyID is missing in the request URL
		/* 		if key == "" {
		   			w.WriteHeader(http.StatusBadRequest)
		   			return
		   		}
		*/value, ok := ch.Msg[key]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, "key not found")
			return
		}
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, value)
	})
}
