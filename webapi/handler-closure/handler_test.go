package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWrapHandler(t *testing.T) {

	respHeaders := RespHeaders{
		{"Header1", "Value1"},
		{"Header2", "Value2"},
		{"Header3", "Value3"},
	}

	mux := http.DefaultServeMux
	mux.Handle("/", WrapHandler(respHeaders))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	mux.ServeHTTP(resp, req)

	actualHeaders := resp.Result().Header
	for _, h := range respHeaders {
		value, ok := actualHeaders[h.key]
		if !ok {
			t.Errorf("header %s not found", h.key)
		}
		if value[0] != h.value {
			t.Errorf("header %q value: get %s, expect %s", h.key, value[0], h.value)
		}
	}

}
