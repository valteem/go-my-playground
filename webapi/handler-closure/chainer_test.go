package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChainer(t *testing.T) {

	hKey, hValue := "Custom-Header", "custom header value"
	body := "response body"
	handlers := []http.HandlerFunc{
		SetHeader(hKey, hValue),
		SetCode(http.StatusOK),
		SetBody(body),
	}

	mux := http.DefaultServeMux
	mux.Handle("/", Chainer(handlers))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	mux.ServeHTTP(resp, req)

	respHeader := resp.Result().Header
	expectedHeaderValue, ok := respHeader[hKey]
	if !ok {
		t.Errorf("header %q not found", hKey)
	} else {
		if expectedHeaderValue[0] != hValue {
			t.Errorf("wrong header %q value: get %s, expect %s", hKey, expectedHeaderValue[0], hValue)
		}
	}

	if respCode := resp.Code; respCode != http.StatusOK {
		t.Errorf("wrong response code: get %d, expect %d", respCode, http.StatusOK)
	}

	if respBody := resp.Body.String(); respBody != body {
		t.Errorf("wrong response body: get %qq, expect %q", respBody, body)
	}
}
