package server

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func startGarbageServer(t *testing.T, requestCount *atomic.Int32) string {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount.Add(1)
		w.WriteHeader(http.StatusInternalServerError)
	}))

	t.Cleanup(server.Close)

	return server.URL

}

func TestGarbageServer(t *testing.T) {

	const numRequests = 100

	count := atomic.Int32{}

	url := startGarbageServer(t, &count)

	for range numRequests {
		resp, err := http.Get(url)
		if err != nil {
			t.Fatalf("failed to make a request: %v", err)
		}
		if status := resp.StatusCode; status != http.StatusInternalServerError {
			t.Errorf("response status: get %d", status)
		}
	}

	if c := count.Load(); c != numRequests {
		t.Errorf("number of requests: get %d, expect %d", c, numRequests)
	}

}
