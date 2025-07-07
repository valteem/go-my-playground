package timeout

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"testing"
)

func TestTimeoutMiddleware(t *testing.T) {

	port := ":8083"

	go runServer(port)
	time.Sleep(100 * time.Millisecond) //allow server some time to start properly

	resp, err := http.Get(fmt.Sprintf("http://localhost%s/quick", port))
	if err != nil {
		t.Fatalf("failed to fetch quick response: %v", err)
	}
	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("quick response status: get %d, expect %d", status, http.StatusOK)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read quick response body: %v", err)
	} else {
		if msg := string(body); msg != quickResponseMsg {
			t.Errorf("quick response message: get %q, expect %q", msg, quickResponseMsg)
		}
	}

	resp, err = http.Get(fmt.Sprintf("http://localhost%s/slow", port))
	if err != nil {
		t.Fatalf("failed to fetch slow response: %v", err)
	}
	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusRequestTimeout {
		t.Errorf("slow response status: get %d, expect %d", status, http.StatusRequestTimeout)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read slow response body: %v", err)
	} else {
		if msg := string(body); msg != timeoutResponseMsg {
			t.Errorf("slow response message: get %q, expect %q", msg, timeoutResponseMsg)
		}
	}
}
