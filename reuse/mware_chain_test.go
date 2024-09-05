package reuse_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/valteem/reuse"
)

const (
	exampleBody        = "example response body"
	exampleHeader      = "Example-Header"
	exampleHeaderValue = "ExampleHeaderValue"
)

func addBody(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(exampleBody))
}

var mwareAddBody reuse.Middleware = func(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addBody(w, r)
		h.ServeHTTP(w, r)
	})
}

func addHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(exampleHeader, exampleHeaderValue)
}

var mwareAddHeader reuse.Middleware = func(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addHeader(w, r)
		h.ServeHTTP(w, r)
	})
}

func addCode(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

var mwareAddCode reuse.Middleware = func(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addCode(w, r)
		h.ServeHTTP(w, r)
	})
}

func doNothing(w http.ResponseWriter, r *http.Request) {
	// do nothing
}

var appDoNothing = http.HandlerFunc(doNothing)

func TestMiddlewareChain(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	mux := http.DefaultServeMux
	// Changing the header map after a call to [ResponseWriter.WriteHeader] (or
	// [ResponseWriter.Write]) has no effect unless the HTTP status code was of the
	// 1xx class or the modified headers are trailers.
	mux.Handle("/", reuse.ChainMiddleware(mwareAddHeader, mwareAddBody, mwareAddCode)(appDoNothing))

	mux.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("status code: get %d, expect %d", resp.Code, http.StatusOK)
	}

	headers := resp.Result().Header
	headerValue, ok := headers[exampleHeader]
	if !ok {
		t.Errorf("response header not found")
	} else {
		if headerValue[0] != exampleHeaderValue {
			t.Errorf("response header value: get %s, expect %s", headerValue[0], exampleHeaderValue)
		}
	}

	responseBodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body: %v", err)
	}
	responseBody := string(responseBodyByte)
	if responseBody != exampleBody {
		t.Errorf("response body: get %s, expect %s", responseBody, exampleBody)
	}

}
