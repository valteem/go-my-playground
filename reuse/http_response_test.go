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
	// Wildcards must be full path segments: they must be preceded by a slash
	// and followed by either a slash or the end of the string
	// https://pkg.go.dev/net/http#ServeMux.ServeHTTP

	/*
		mux.HandleFunc("/?model={model_id}&qty={qty}", func(w http.ResponseWriter, r *http.Request) {
			modelId := r.PathValue("model_id")
			qty := r.PathValue("qty")
			io.WriteString(w, "model: "+modelId+", quantity: "+qty)
		})
	*/
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

	tests := []struct {
		path   string
		output string
	}{
		{path: "/model/myModel/qty/1234",
			output: "model: myModel, quantity: 1234",
		},

		{path: "/?model=myModel&qty=1234",
			output: "404 page not found\n",
		},
	}

	for _, tc := range tests {
		req, err = http.NewRequest("GET", tc.path, nil)
		if err != nil {
			t.Errorf("Creating 'GET %s' request failed!", tc.path)
		}

		m.ServeHTTP(respRec, req)

		buf := new(strings.Builder)
		_, e := io.Copy(buf, respRec.Body)
		if e != nil {
			t.Errorf("Error reading response body: %q", e)
		}

		if output := buf.String(); output != tc.output {
			t.Errorf("Response body: get %s, expect %s", output, tc.output)
		}
	}

}
