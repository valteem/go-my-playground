package main

import (
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
)

const (
	srvPort = ":3001"
)

func runServer(sm *scs.SessionManager) {

	mux := http.NewServeMux()
	mux.Handle("/create", createMessage(sm))
	mux.Handle("/receive", receiveMessage(sm))

	err := http.ListenAndServe(srvPort, sm.LoadAndSave(mux))
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
func TestSessions(t *testing.T) {

	sm := scs.New()
	sm.Lifetime = 24 * time.Hour

	go runServer(sm)

	jar, err := cookiejar.New(nil) // no option
	if err != nil {
		t.Fatalf("failed to create cookie jar: %v", err)
	}
	client := http.Client{
		Jar: jar,
	}

	urlCreate := "http://localhost" + srvPort + "/create"
	reqCreate, err := http.NewRequest(http.MethodPost, urlCreate, nil)
	if err != nil {
		t.Fatalf("failed to create a /create request: %v", err)
	}
	_, err = client.Do(reqCreate)
	if err != nil {
		t.Fatalf("failed to get a response to /create request: %v", err)
	}

	urlReceive := "http://localhost" + srvPort + "/receive"
	reqReceive, err := http.NewRequest(http.MethodGet, urlReceive, nil)
	if err != nil {
		t.Fatalf("failed to create a /receive request: %v", err)
	}
	resp, err := client.Do(reqReceive)
	if err != nil {
		t.Fatalf("failed to get a response to /receive request: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	bodyStr := string(bodyBytes)
	if bodyStr != msgBody {
		t.Errorf("response message: get %q, expect %q", bodyStr, msgBody)
	}

}
