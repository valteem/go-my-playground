package examine

import (
	"net/http"
)

type mockResponseWriter struct{}

func (w *mockResponseWriter) Header() http.Header {
	return http.Header{}
}

func (w *mockResponseWriter) Write(b []byte) (int, error) {
	return 0, nil
}

func (w *mockResponseWriter) WriteHeader(statusCode int) {}
