// https://stackoverflow.com/a/55527412
package main

import (
	"net/http"
	"testing"
	"time"

	"webapi/wrap-auth/auth"
	"webapi/wrap-auth/handlers"
)

const (
	headerWrapAuth = "X-Wrap-Auth"
)

var (
	incorrectToken = "not_authorized"
	correctToken   = "authorized"
)

func runServer() {

	mux := http.NewServeMux()
	mux.Handle("/hello", http.HandlerFunc(handlers.HandleHello))
	mux.Handle("/info", auth.GrantAccess(handlers.HandleInfo))

	http.ListenAndServe(":3001", mux)

}

func TestWrapAuth(t *testing.T) {

	go runServer()
	time.Sleep(100 * time.Millisecond) // allow server some time to start

	client := http.Client{}

	// incorrect token - public path
	req, err := http.NewRequest(http.MethodGet, "http://localhost:3001/hello", nil)
	if err != nil {
		t.Fatalf("failed to create a request: %v", err)
	}
	req.Header.Add("Authorization", incorrectToken)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to get response: %v", err)
	}
	defer resp.Body.Close()
	if msg := resp.Header[headerWrapAuth][0]; msg != "hello" {
		t.Errorf("get %q, expect %q", msg, "hello")
	}

	// correct token - public path
	{
		req, err = http.NewRequest(http.MethodGet, "http://localhost:3001/hello", nil)
		if err != nil {
			t.Fatalf("failed to create a request: %v", err)
		}
		req.Header.Add("Authorization", correctToken)
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("failed to get response: %v", err)
		}
		defer resp.Body.Close()
		if msg := resp.Header[headerWrapAuth][0]; msg != "hello" {
			t.Errorf("get %q, expect %q", msg, "hello")
		}
	}

	// correct token - secured path
	{
		req, err = http.NewRequest(http.MethodGet, "http://localhost:3001/info", nil)
		if err != nil {
			t.Fatalf("failed to create a request: %v", err)
		}
		req.Header.Add("Authorization", correctToken)
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("failed to get response: %v", err)
		}
		defer resp.Body.Close()
		if msg := resp.Header[headerWrapAuth][0]; msg != "info" {
			t.Errorf("get %q, expect %q", msg, "info")
		}
	}

	// incorrect token - secured path
	{
		req, err = http.NewRequest(http.MethodGet, "http://localhost:3001/info", nil)
		if err != nil {
			t.Fatalf("failed to create a request: %v", err)
		}
		req.Header.Add("Authorization", incorrectToken)
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("failed to get response: %v", err)
		}
		defer resp.Body.Close()
		msgs, ok := resp.Header[headerWrapAuth]
		if ok {
			if msg := msgs[0]; msg == "info" {
				t.Errorf("get %q, expect %q", msg, "")
			}
		}
		if status := resp.Status; status != "403 Forbidden" {
			t.Errorf("response status: get %q, expect %q", status, http.StatusForbidden)
		}
	}

}
