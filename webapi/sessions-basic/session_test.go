package main

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("super-secret-key") // Key to encrypt cookie-session
	store = sessions.NewCookieStore(key)
)

func createSessionHandler(w http.ResponseWriter, r *http.Request) {

	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, _ := store.Get(r, "session-name")

	// Cookies flagged as secure should only be sent over a secure link
	// https://stackoverflow.com/a/79161740/16648033
	// localhost should be an exception, but it does not work:
	// https://github.com/golang/go/issues/60997
	//https://github.com/grafana/k6/issues/3457
	session.Options.Secure = false

	// Set some session values.
	session.Values["weather"] = "fine"
	session.Values["winter"] = "cold"

	// Save it before we write to the response/return from the handler.
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func useSessionHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session-name") // retrieve existing session

	weather := session.Values["weather"].(string)
	winter := session.Values["winter"].(string)

	w.Write([]byte("weather is " + weather + " and winter is " + winter))

}

func runServer() {

	mux := http.NewServeMux()

	mux.Handle("/create", http.HandlerFunc(createSessionHandler))
	mux.Handle("/use", http.HandlerFunc(useSessionHandler))

	http.ListenAndServe(":3001", mux)

}

func TestSessionBasic(t *testing.T) {

	jar, err := cookiejar.New(nil) //no options
	if err != nil {
		t.Fatalf("failed to create new coockie jar: %v", err)

	}

	client := &http.Client{
		Jar: jar,
	}

	go runServer()

	time.Sleep(1 * time.Second)

	respCreate, err := client.Get("http://localhost:3001/create")
	if err != nil {
		t.Fatalf("failed to get response for create: %v", err)
	}
	defer respCreate.Body.Close()

	respUse, err := client.Get("http://localhost:3001/use")
	if err != nil {
		t.Fatalf("failed to get response for use: %v", err)
	}
	defer respUse.Body.Close()

	body, err := io.ReadAll(respUse.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	if actual, expected := string(body), "weather is fine and winter is cold"; actual != expected {
		t.Errorf("response to \"use\" body:\nget\n%q\nexpect\n%q", actual, expected)
	}

}
