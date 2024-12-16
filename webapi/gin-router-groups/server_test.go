package main

import (
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestRouterGroups(t *testing.T) {

	tests := []struct {
		path     string
		response string
	}{
		{"/v1/items", "items"},
		{"/v1/items/descriptions", "descriptions"},
		{"/v1/items/images", "images"},
		{"/v1/stores", "stores"},
		{"/v1/stores/locations", "locations"},
		{"/v2/items", "items"},
		{"/v2/items/descriptions", "descriptions"},
		{"/v2/items/images", "images"},
		{"/v2/stores", "stores"},
		{"/v2/stores/locations", "locations"},
	}

	e := gin.Default()

	go runServer(e, ":3001")

	time.Sleep(100 * time.Millisecond) // allow server some time to start properly

	client := http.Client{}

	for _, tc := range tests {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:3001"+tc.path, nil)
		if err != nil {
			t.Fatalf("failed to create request %q: %v", tc.path, err)
		}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("failed to fetch response for %q: %v", tc.path, err)
		}
		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read %q response body: %v", tc.path, err)
		}
		s := string(b)
		if s != tc.response {
			t.Errorf("%q: get %q, expect %q", tc.path, s, tc.response)
		}
	}

}
