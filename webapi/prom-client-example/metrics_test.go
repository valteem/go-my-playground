package main

import (
	"io"
	"net/http"
	"testing"
	"time"
)

func TestMetrics(t *testing.T) {

	go run()

	// allow endpoint some time to start properly
	time.Sleep(1 * time.Second)

	req, err := http.NewRequest(http.MethodGet, "http://localhost:3001/metrics", nil)
	if err != nil {
		t.Fatalf("failed to create new request: %v", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to fetch response: %v", err)
	}
	defer resp.Body.Close()

	b := make([]byte, 1024)
	n, err := resp.Body.Read(b)
	if err != nil && err != io.EOF {
		t.Fatalf("failed to read response body: %v", err)
	}

	b = b[:n]

	if len(b) == 0 {
		t.Errorf("empty metrics response")
	}

}
