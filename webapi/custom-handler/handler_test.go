package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCustomHandler(t *testing.T) {

	mux := New()
	mux.AddMessage("key1", "value1")
	mux.AddMessage("key2", "value2")

	mux.Handle("/key/{keyID}", mux.handleMsg())

	tests := []struct {
		input    string
		respCode int
		respMsg  string
	}{
		{"/key/key1", http.StatusOK, "value1"},
		{"/key/key2", http.StatusOK, "value2"},
		{"/key/key3", http.StatusNotFound, "key not found"},
		{"/key", http.StatusNotFound, "404 page not found\n"},
	}

	for _, tc := range tests {
		req := httptest.NewRequest(http.MethodGet, tc.input, nil)
		resp := httptest.NewRecorder()
		mux.ServeHTTP(resp, req)
		respCode := resp.Code
		respMsg := resp.Body.String()
		if respCode != tc.respCode || respMsg != tc.respMsg {
			t.Errorf("%q: get(%d, %q), expect (%d, %q)", tc.input, respCode, respMsg, tc.respCode, tc.respMsg)
		}
	}

}
