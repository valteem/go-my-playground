package ensure

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"time"

	"testing"

	_ "webapi/jwtmware/env"
)

func TestAuth(t *testing.T) {

	_ = Start()

	time.Sleep(100 * time.Millisecond) // allow server some time to start properly

	// First log in and receive access token
	ut := userCred{Username: cfg.UserName, Password: cfg.Password}
	utJson, err := json.Marshal(ut)
	if err != nil {
		t.Fatalf("failed to marshal user credentials: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:3001/login", bytes.NewReader([]byte(utJson)))
	if err != nil {
		t.Fatalf("failed to create authorization request: %v", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to fetch server response to authorization request: %v", err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read authorization response body: %v", err)
	}
	if len(b) == 0 {
		t.Errorf("empty auth token")
	}

	at := authToken{}
	// if err := json.NewDecoder(resp.Body).Decode(&at); err != nil {
	if err := json.Unmarshal(b, &at); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	if len(at.AccessToken) == 0 {
		t.Fatalf("access token empty")
	}

	// Next seng ping request and get pong response
	req, err = http.NewRequest(http.MethodGet, "http://localhost:3001/ping", nil)
	if err != nil {
		t.Fatalf("failed to create ping request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", at.AccessToken))

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("failed to fetch server response to ping request: %v", err)
	}
	defer resp.Body.Close()

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read ping response body: %v", err)
	}

	if actual, expected := string(b), "pong"; actual != expected {
		t.Errorf("ping response: get %q, expect %q", actual, expected)
	}

}

func TestHandlers(t *testing.T) {

	loadConfig(context.Background())

	mux := http.NewServeMux()

	s := &Server{
		Server: http.Server{
			Handler: mux,
		},
	}

	mux.HandleFunc("/login", s.handleLogin)
	mux.Handle("/ping", NewEnsureAuth(s.handlePing))

	ut := userCred{Username: cfg.UserName, Password: cfg.Password}
	utJson, err := json.Marshal(ut)
	if err != nil {
		t.Fatalf("failed to marshal user credentials: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewReader([]byte(utJson)))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp := httptest.NewRecorder()

	mux.ServeHTTP(resp, req)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	at := authToken{}
	if err := json.Unmarshal(respBody, &at); err != nil {
		t.Fatalf("failed to unmarshal response body to auth token: %v", err)
	}

	req, err = http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		t.Fatalf("failed to create GET request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", at.AccessToken))

	mux.ServeHTTP(resp, req)

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body of ping response: %v", err)
	}

	if actual, expected := string(respBody), "pong"; actual != expected {
		t.Errorf("ping response: get %q, expect %q", actual, expected)
	}

}
