package reuse_test

import (
	"context"
	"io"
	"log"
	"net/http"
	"testing"
	"time"
)

func sendRequestgetAndCheckResponse(t *testing.T, method string, URL string) string {
	req, err := http.NewRequest(method, URL, nil)
	if err != nil {
		t.Fatalf("Error setting up new request: %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Error running http client: %v", err)
	}
	// The caller must close the response body when finished with it
	// https://pkg.go.dev/net/http#:~:text=The%20caller%20must%20close%20the%20response%20body%20when%20finished%20with%20it
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}
	s := string(b[:])
	if len(s) == 0 {
		t.Errorf("Empty response body")
	}
	return s
}
func TestRequestToExternalServer(t *testing.T) {
	sendRequestgetAndCheckResponse(t, "GET", "https://google.com/")
}

func TestLocalHttpServer(t *testing.T) {
	portNum := "3000"
	responseMessage := "Local response"
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(responseMessage))
	}
	srv := &http.Server{
		Addr:    ":" + portNum,
		Handler: http.HandlerFunc(handleFunc),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	time.Sleep(1 * time.Second) // allow some time for the server to start
	respMsg := sendRequestgetAndCheckResponse(t, "GET", "http://localhost:"+portNum)
	if respMsg != responseMessage {
		t.Errorf("response message: get %q, expect %q", respMsg, responseMessage)
	}
	srv.Shutdown(context.Background())
}

type stubHandler struct{}

func (s stubHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func TestPatternConflict(t *testing.T) {

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("two conflicting pattern should panic")
		}
	}()

	mux := http.NewServeMux()
	pattern := "/some_pattern"
	mux.Handle(pattern, stubHandler{})
	mux.Handle(pattern, stubHandler{})

}
