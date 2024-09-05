package reuse

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func ChainMiddleware(mware ...Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		next := h
		for i := len(mware) - 1; i >= 0; i-- {
			next = mware[i](next)
		}
		return next
	}
}
