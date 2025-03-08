package main

import (
	"net/http"
	"testing"
	"time"
)

func TestChainHandlers(t *testing.T) {

	headersExpected := map[string]string{
		"Custom-Header-A": "custom-header-a-value",
		"Custom-Header-B": "custom-header-b-value",
		"Custom-Header-C": "custom-header-c-value",
	}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	})

	go runServer(h)

	// allow server some time to start properly
	time.Sleep(100 * time.Millisecond)

	req, err := http.NewRequest(http.MethodGet, "http://localhost:3001/headers", nil)
	if err != nil {
		t.Fatalf("failed to create new http request: %v", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to fetch http response: %v", err)
	}
	defer resp.Body.Close()

	for k, v := range headersExpected {
		a, ok := resp.Header[k]
		if !ok {
			t.Errorf("header %q is missing", k)
		}
		if a[0] != v {
			t.Errorf("header %q value: get %q, expect %q", k, a[0], v)
		}
	}

}
