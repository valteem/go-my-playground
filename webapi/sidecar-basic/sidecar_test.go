package main

import (
	"io"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestSidecarBasic(t *testing.T) {

	go runPrimary()

	go runProxy()

	time.Sleep(1 * time.Second)

	resp, err := http.Get("http://127.0.0.1:3002/")
	if err != nil {
		log.Fatalf("failed to get proxy response: %v", err)
	}

	proxyHeaderActual, ok := resp.Header[http.CanonicalHeaderKey("x-sidecar-status")]
	if !ok {
		t.Fatalf("missing proxy header")
	}
	if proxyHeaderExpected := "sidecar-running"; proxyHeaderActual[0] != proxyHeaderExpected {
		t.Errorf("proxy header: get\n%q\nexpect\n%q\n", proxyHeaderActual[0], proxyHeaderExpected)
	}

	primaryHeaderActual, ok := resp.Header[http.CanonicalHeaderKey("x-primary-status")]
	if !ok {
		t.Fatalf("missing primary header")
	}
	if primaryHeaderExpected := "primary-running"; primaryHeaderActual[0] != primaryHeaderExpected {
		t.Errorf("primary header: get\n%q\nexpect\n%q\n", primaryHeaderActual[0], primaryHeaderExpected)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	defer resp.Body.Close()
	if msgActual, msgExpected := string(body), "primary service response"; msgActual != msgExpected {
		t.Errorf("service message:\nget\n%q\nexpect\n%q", msgActual, msgExpected)
	}

}
