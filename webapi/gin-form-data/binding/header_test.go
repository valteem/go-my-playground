package binding

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func runServer(port string) {
	e := gin.Default()
	e.POST("/store", HandleNewStore)
	err := e.Run(port)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
		return
	}
}

func TestShouldBindHeader(t *testing.T) {

	tests := []struct {
		header   string
		value    string
		payload  string
		respCode int
	}{
		{
			header:   "storeid",
			value:    "1001",
			payload:  `{"location":"somewhere", "square":10000}`,
			respCode: http.StatusOK,
		},
		{
			header:   "itemid", // invalid header ...
			value:    "1001",
			payload:  `{"location":"somewhere", "square":10000}`,
			respCode: http.StatusBadRequest, // ... should not bind
		},
	}

	go runServer(":3001")

	time.Sleep(100 * time.Millisecond) // allow server some time to start properly

	client := http.Client{}

	for _, tc := range tests {

		req, err := http.NewRequest(http.MethodPost, "http://localhost:3001/store", strings.NewReader(tc.payload))
		if err != nil {
			t.Fatalf("failed to create a request: %v", err)
		}
		req.Header.Add(tc.header, tc.value)

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("failed to fetch response: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != tc.respCode {
			t.Fatalf("response code: get %d, expect %d", resp.StatusCode, tc.respCode)
		}

		if resp.StatusCode == http.StatusOK {

			var store *Store
			// buf := make([]byte, 4)
			// resp.Body.Read(buf)

			err = json.NewDecoder(resp.Body).Decode(&store)
			if err != nil {
				t.Fatalf("failed to decode response body to JSON: %v", err)
			}

			if store.Location != "somewhere" || store.Square != 10000 {
				t.Errorf("location/square: get (%s/%d), expect (%s/%d)", store.Location, store.Square, "somewhere", 10000)
			}

		}

	}

}

func TestRequestHeaderBind(t *testing.T) {

	tests := []struct {
		input  http.Header
		output *Person
	}{
		{
			input:  http.Header{"Name": []string{"SomeName"}, "Age": []string{"42"}},
			output: &Person{Name: "SomeName", Age: 42},
		},
		{
			input:  http.Header{"Name": []string{"SomeName"}, "Address": []string{"SomeAddress"}},
			output: &Person{Name: "SomeName", Age: 0},
		},
	}

	for _, tc := range tests {
		c := gin.Context{}
		req := &http.Request{}
		req.Header = tc.input
		c.Request = req
		output := &Person{}
		c.ShouldBindHeader(output)
		if !reflect.DeepEqual(output, tc.output) {
			t.Errorf("get\n%v\nexpect\n%v\n", output, tc.output)
		}
	}

}
