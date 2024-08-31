package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRoutesPerson(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/person?name=somename&age=33", nil)
	res := httptest.NewRecorder()

	g := gin.Default()
	g.GET("/person", GetPerson)

	g.ServeHTTP(res, req)

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error reading output: %v", err)
	}
	JSONactual := string(bytes)
	JSONexpected := `{"age":33,"name":"somename"}`
	if JSONactual != JSONexpected {
		t.Errorf("JSON output: get %s, expect %s", JSONactual, JSONexpected)
	}

}
