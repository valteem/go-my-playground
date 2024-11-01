package main

import (
	"io"
	"net/http"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {

	go runHandler()

	time.Sleep(1 * time.Second)

	resp, err := http.Get("http://127.0.0.1:3001")

	if err != nil {
		t.Fatalf("failed to get a response: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	actual := string(body)
	if expected := "service is up and running"; actual != expected {
		t.Errorf("server response:\nget\n%q\nexpect\n%q", actual, expected)
	}

}
