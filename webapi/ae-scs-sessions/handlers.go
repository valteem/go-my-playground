package main

import (
	"io"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

const (
	msgKey  = "message"
	msgBody = "some custom message"
)

func createMessage(sm *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sm.Put(r.Context(), msgKey, msgBody)
	}
}

func receiveMessage(sm *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := sm.GetString(r.Context(), msgKey)
		io.WriteString(w, msg)
	}
}
