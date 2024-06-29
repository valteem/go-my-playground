package reuse_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var m *http.ServeMux
var req *http.Request
var err error
var respRec *httptest.ResponseRecorder

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/model/{model_id}/qty/{qty}", func(w http.ResponseWriter, r *http.Request) {
		modelId := r.PathValue("model_id")
		qty := r.PathValue("qty")
		io.WriteString(w, "model: "+modelId+", quantity: "+qty)
	})
}

func setup() {
	//mux router with added question routes
	m = http.NewServeMux()
	AddRoutes(m)

	//The response recorder used to record HTTP responses
	respRec = httptest.NewRecorder()
}

func TestResponseBody(t *testing.T) {
	setup()

	path := "/model/myModel/qty/1234"
	req, err = http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal("Creating 'GET /model/myModel/qty/1234' request failed!")
	}

	m.ServeHTTP(respRec, req)

	buf := new(strings.Builder)
	_, e := io.Copy(buf, respRec.Body)
	if e != nil {
		t.Errorf("Error reading response body: %q", e)
	}

	expectedRespBody := "model: myModel, quantity: 1234"
	if respBody := buf.String(); respBody != expectedRespBody {
		t.Errorf("Response body: get %s, expect %s", respBody, expectedRespBody)
	}
}
