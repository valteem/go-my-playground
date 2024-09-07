package main

import (
	"net/http"
)

type RespHeader struct {
	key   string
	value string
}

type RespHeaders []RespHeader

func WrapHandler(rh RespHeaders) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		for _, h := range rh {
			w.Header().Set(h.key, h.value)
		}
	}
	return fn
}
