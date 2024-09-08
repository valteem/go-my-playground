package main

import (
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

	JSONactual := res.Body.String()
	JSONexpected := `{"age":33,"name":"somename"}`
	if JSONactual != JSONexpected {
		t.Errorf("JSON output: get %s, expect %s", JSONactual, JSONexpected)
	}

}
