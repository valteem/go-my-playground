package reuse_test

import (
	"io"
	"net/http"
	"testing"
)

func TestNewRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "https://google.com/", nil)
	if err != nil {
		t.Errorf("Error setting up new request: %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Error running http client: %v", err)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	if s := string(b); len(s) == 0 {
		t.Errorf("Empty response body")
	}
}
