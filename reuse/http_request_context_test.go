package reuse_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// type ctxKey uint8

const (
	ctxFlag   ctxKey = 1
	valueFlag string = "custom flag"
)

func TestRequestContext(t *testing.T) {

	flagCtx := context.WithValue(context.Background(), ctxFlag, valueFlag)

	/*
		server := &http.Server{
			Addr: "8080",
			BaseContext: func(net.Listener) context.Context {
				return flagCtx
			},
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				flag := ctx.Value(ctxFlag).(string)
				w.Write([]byte(flag))
			}),
		}
	*/

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	r := req.WithContext(flagCtx) // emulate server.BaseContext()
	resp := httptest.NewRecorder()

	handler := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		flag := ctx.Value(ctxFlag).(string)
		w.Write([]byte(flag))
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handler))

	mux.ServeHTTP(resp, r)

	output := resp.Body.String()
	if output != valueFlag {
		t.Errorf("response body: get %q, expect %q", output, valueFlag)
	}

}
