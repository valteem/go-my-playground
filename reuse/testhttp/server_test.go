package testhttp

import (
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerURL(t *testing.T) {

	srvAddr := "127.0.0.1:3001"

	l, err := net.Listen("tcp", srvAddr)
	if err != nil {
		t.Fatalf("failed to create tcp listener: %v", err)
	}

	handler := http.NewServeMux()

	ts := httptest.NewUnstartedServer(handler)
	ts.Listener.Close()
	ts.Listener = l

	ts.Start()
	defer ts.Close()

	if actual, expected := ts.URL, "http://"+srvAddr; actual != expected {
		t.Errorf("test server URL: get %q, expect %q", actual, expected)
	}

}

func TestServerBaseURL(t *testing.T) {

	handler := http.NewServeMux()
	path := "/test"
	headerKey, headerValue := "Custom-Header", "custom header value"
	handler.Handle(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(headerKey, headerValue)
	}))

	ts := httptest.NewServer(handler)

	baseURL := ts.URL

	req, _ := http.NewRequest(http.MethodGet, baseURL+path, nil)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to fetch server response: %v", err)
	}
	defer resp.Body.Close()

	respHeaderValue := resp.Header[headerKey][0]
	if respHeaderValue != headerValue {
		t.Errorf("custom header %q value: get %q, expect %q", headerKey, respHeaderValue[0], headerValue)
	}

}
