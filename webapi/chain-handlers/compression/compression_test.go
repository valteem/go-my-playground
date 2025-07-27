package compression

import (
	"bytes"
	"compress/gzip"
	"net/http"
	"net/http/httptest"

	"testing"
)

func TestGzipResponse(t *testing.T) {

	msg := "response from server"

	mux := http.NewServeMux()
	mux.Handle("GET /message", WithCompressionGzip(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(msg))
	})))

	srv := httptest.NewServer(mux)
	defer srv.Close()

	client := &http.Client{
		Transport: &http.Transport{},
	}

	req, err := http.NewRequest(http.MethodGet, srv.URL+"/message", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to fetch server response: %v", err)
	}
	defer resp.Body.Close()

	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		t.Fatalf("failed to create new gzip reader: %v", err)
	}

	var buf bytes.Buffer
	_, err = buf.ReadFrom(gzipReader)
	if err != nil {
		t.Fatalf("failed to fetch gzip reader output: %v", err)
	}

	if actual, expected := buf.String(), msg; actual != expected {
		t.Errorf("server response: get %q, expect %q", actual, expected)
	}

}
