package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHandlers(t *testing.T) {

	router := mux.NewRouter()
	router.HandleFunc("/articles/{id:[0-9]+}", handleNumbers)
	router.HandleFunc("/articles/{id:[a-zA-Z]+}", handleLetters)
	router.HandleFunc("/articles/{id:[a-zA-Z0-9]+}", handleNumbersAndLetters)
	router.HandleFunc("/articles/{id:[^a-zA-Z0-9]+}", handleNoNumbersOrLetters)

	tests := []struct {
		input  string
		output string
	}{
		{"/articles/1234", "numbers: 1234"},
		{"/articles/abcUVW", "letters: abcUVW"},
		{"/articles/pqr456", "numbers and letters: pqr456"},
		// https://www.w3schools.com/tags/ref_urlencode.ASP
		{"/articles/%21%23%24%25", "something strange: !#$%"},
	}

	for _, tc := range tests {
		req := httptest.NewRequest(http.MethodGet, tc.input, nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		output := resp.Body.String()
		if output != tc.output {
			t.Errorf("for input %s: get %s, expect %s", tc.input, output, tc.output)
		}
	}

}
