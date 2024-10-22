package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/api/query?partnum=1&store=central", nil)
	resp := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/query", handleBothGetAndPost)

	mux.ServeHTTP(resp, req)

	responseBody, err := io.ReadAll(resp.Result().Body)
	if err != nil {
		t.Fatalf("error reading response body: %v", err)
	}

	var responseActual Response
	err = json.Unmarshal(responseBody, &responseActual)
	if err != nil {
		t.Fatalf("error decoding response body: %v", err)
	}

	responseExpected := Response{PartNum: "1", Store: "central", Qty: 1}

	if !reflect.DeepEqual(responseActual, responseExpected) {
		t.Errorf("response:\nget%v\nexpect\n%v", responseActual, responseExpected)
	}

}

func TestPost(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/api/query", strings.NewReader("partnum=1&store=central"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/query", handleBothGetAndPost)

	mux.ServeHTTP(resp, req)

	responseBody, err := io.ReadAll(resp.Result().Body)
	if err != nil {
		t.Fatalf("error reading response body: %v", err)
	}

	var responseActual Response
	err = json.Unmarshal(responseBody, &responseActual)
	if err != nil {
		t.Fatalf("error decoding response body: %v", err)
	}

	responseExpected := Response{PartNum: "1", Store: "central", Qty: 1}

	if !reflect.DeepEqual(responseActual, responseExpected) {
		t.Errorf("response:\nget%v\nexpect\n%v", responseActual, responseExpected)
	}

}
