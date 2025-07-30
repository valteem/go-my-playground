package trace

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"
)

var (
	errMsg                    = "error processing request"
	panicHandler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(errMsg)
	})
)

func TestWithStackTrace(t *testing.T) {

	mux := http.NewServeMux()
	mux.Handle("/path", WithStackTrace(panicHandler))

	srv := httptest.NewServer(mux)
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/path")
	if err != nil {
		t.Fatalf("failed to fetch response: %v", err)
	}
	defer resp.Body.Close()

	if actual, expected := resp.StatusCode, http.StatusBadRequest; actual != expected {
		t.Errorf("response status code: get %s, expect %s", http.StatusText(actual), http.StatusText(expected))
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	output := string(buf)
	if !strings.Contains(output, errMsg) {
		t.Errorf("response body does not contain handler error message")
	}

}
