package routes

import (
	"net/http"
	"testing"
)

var (
	routeFlag      int
	mux            http.ServeMux
	responseWriter *mockResponseWriter
)

func wrapSomeHandler(w http.ResponseWriter, r *http.Request) {
	routeFlag = 1
	someHandler(w, r)
}

func wrapSomeOtherHandler(w http.ResponseWriter, r *http.Request) {
	routeFlag = 2
	someOtherHandler(w, r)
}

func setRoutes() {
	mux.Handle("/someURLpath", http.HandlerFunc(wrapSomeHandler))
	mux.Handle("/someOtherURLpath", http.HandlerFunc(wrapSomeOtherHandler))
}

func TestRoutes(t *testing.T) {

	tests := []struct {
		path string
		flag int
	}{
		{"/someURLpath", 1},
		{"/someOtherURLpath", 2},
	}

	setRoutes()

	for _, tc := range tests {
		req, _ := http.NewRequest(http.MethodGet, "https://example.com"+tc.path, nil)
		mux.ServeHTTP(responseWriter, req)
		if routeFlag != tc.flag {
			t.Errorf("route flag: get %d, expect %d", routeFlag, tc.flag)
		}
	}
}
