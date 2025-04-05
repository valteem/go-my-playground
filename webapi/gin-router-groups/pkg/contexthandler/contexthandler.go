package contexthandler

import (
	"context"
	"net/http"
)

type ctxKey int

const (
	chCtxKey  ctxKey = iota
	someValue string = "some value"
)

type ContextHandler struct {
	h http.Handler
}

func setValue(r *http.Request, val string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), chCtxKey, val))
}

func (ch ContextHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	r := setValue(req, someValue)

	ch.h.ServeHTTP(rw, r)

}
