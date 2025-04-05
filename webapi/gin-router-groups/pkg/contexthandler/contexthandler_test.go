package contexthandler

import (
	"net/http"
	"testing"
	"time"
)

func TestContextHandler(t *testing.T) {

	customHeader := "Added-Context-Value" // canonical format

	ch := ContextHandler{http.DefaultServeMux}

	mux := http.DefaultServeMux
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		value := (r.Context().Value(chCtxKey)).(string)
		w.Header().Set(customHeader, value)
	}))

	go http.ListenAndServe(":3001", ch)

	// allow server some time to start properly
	time.Sleep(100 * time.Millisecond)

	req, err := http.NewRequest(http.MethodGet, "http://localhost:3001", nil)
	if err != nil {
		t.Fatalf("failed to create a request: %v", err)
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to fetch server response: %v", err)
	}

	header := resp.Header[customHeader]
	value := header[0]
	if value != someValue {
		t.Errorf("get %q, expect %q", value, someValue)
	}

}
