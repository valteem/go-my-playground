package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

const (
	serverResponse = "some response that is first gzipped at server and than unzipped at client"
)

func runServer() {

	mux := http.NewServeMux()

	mux.Handle("/gzip", http.HandlerFunc(HandleGZipResponse(serverResponse)))

	http.ListenAndServe(":3001", mux)

}

func TestGZipResponse(t *testing.T) {

	client := &http.Client{}

	go runServer()
	// allow server some time to start properly
	time.Sleep(1 * time.Second)

	resp, err := client.Get("http://localhost:3001/gzip")
	if err != nil {
		t.Fatalf("failed to fetch server response: %v", err)
	}
	defer resp.Body.Close()

	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		t.Fatalf("failed to create new gzip reader: %v", err)
	}

	outputBytes, err := io.ReadAll(gzipReader)
	if err != nil {
		t.Fatalf("failed to read unzipped response body: %v", err)
	}

	if actual, expected :=
		string(outputBytes),
		fmt.Sprintf("%q", serverResponse); actual != expected {
		t.Errorf("unzip server response:\nget\n%s\nexpect\n%s\n", actual, expected)
	}

}
