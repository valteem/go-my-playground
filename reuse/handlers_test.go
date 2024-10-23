package reuse_test

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/valteem/reuse"
)

const panicMsg = "test handler - panic"

func TestLogStackTraceOnHandlerPanic(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	mux := http.NewServeMux()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(panicMsg))
		panic(panicMsg)
	})

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux.Handle("/", reuse.LogStackTraceOnHandlerPanic(handler, logger))

	mux.ServeHTTP(resp, req)

	// test does not get there - panic
	if output := resp.Body.String(); output != panicMsg {
		t.Errorf("response body: get %s, expect %s", output, panicMsg)
	}

}
