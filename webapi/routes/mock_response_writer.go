package routes

import "net/http"

const mockResponse = "mock response"

type mockResponseWriter struct {
	code int
}

func (m *mockResponseWriter) Header() http.Header {
	return http.Header{}
}

func (m *mockResponseWriter) Write([]byte) (int, error) {
	return len(mockResponse), nil
}

func (m *mockResponseWriter) WriteHeader(statusCode int) {
	m.code = statusCode
}
