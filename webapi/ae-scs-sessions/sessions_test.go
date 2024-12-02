package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/jackc/pgx/v5/pgxpool"
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

func runClient() (string, error) {

	jar, err := cookiejar.New(nil) // no option
	if err != nil {
		return "", fmt.Errorf("failed to create cookie jar: %v", err)
	}
	client := http.Client{
		Jar: jar,
	}

	urlCreate := "http://localhost" + srvPort + "/create"
	reqCreate, err := http.NewRequest(http.MethodPost, urlCreate, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create a /create request: %v", err)
	}
	_, err = client.Do(reqCreate)
	if err != nil {
		return "", fmt.Errorf("failed to get a response to /create request: %v", err)
	}

	urlReceive := "http://localhost" + srvPort + "/receive"
	reqReceive, err := http.NewRequest(http.MethodGet, urlReceive, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create a /receive request: %v", err)
	}
	resp, err := client.Do(reqReceive)
	if err != nil {
		return "", fmt.Errorf("failed to get a response to /receive request: %v", err)
	}
	defer resp.Body.Close()

	msgBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}
	msgStr := string(msgBytes)

	return msgStr, nil

}
func TestSessions(t *testing.T) {

	sm := scs.New()
	sm.Lifetime = 24 * time.Hour

	go runServer(sm)

	msg, err := runClient()
	if err != nil {
		t.Fatalf("failed to run client part: %v", err)
	}
	if msg != msgBody {
		t.Errorf("response message: get %q, expect %q", msg, msgBody)
	}

}

func TestSessionsWithPGXStore(t *testing.T) {

	connStr := os.Getenv("PGXSTORE_CONN")
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	sm := scs.New()
	sm.Store = pgxstore.New(pool)
	sm.Lifetime = 24 * time.Hour

	go runServer(sm)

	time.Sleep(1 * time.Second) // allow server some time to start

	msg, err := runClient()
	if err != nil {
		t.Fatalf("failed to run client part: %v", err)
	}
	if msg != msgBody {
		t.Errorf("response message: get %q, expect %q", msg, msgBody)
	}

}
