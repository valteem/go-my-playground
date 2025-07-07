package tree

import (
	"fmt"
	"net/http"
	"time"

	"testing"
)

func TestRouting(t *testing.T) {

	port := ":8083"

	go runServer(port)
	time.Sleep(100 * time.Millisecond) // allow server some time to start properly

	routes := []string{
		"",
		"/",
		"/account",
		"/account/",
		"/account/active",
		"/account/active/",
	}

	for _, route := range routes {
		_, err := http.Get(fmt.Sprintf("http://127.0.0.1%s%s", port, route))
		if err != nil {
			t.Errorf("failed to fetch response for route %q: %v", route, err)
			continue
		}
	}

}
