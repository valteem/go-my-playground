package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

var (
	inputJSON = `{"desc":"someNewItem"}`
	fmtOutput = `{"id":%d,"desc":"someNewItem"}`
)

func TestCreateItem(t *testing.T) {

	e := echo.New()
	for i := 1; i < 11; i++ {
		req := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(inputJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) // Content-Type: application/json
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if e := CreateItem(c); e != nil {
			t.Fatalf("error creating item: %v", e)
		}
		if rec.Code != http.StatusCreated {
			t.Errorf("expect response code %d, get %d", http.StatusCreated, rec.Code)
		}
		// https://github.com/labstack/echo/discussions/2450
		if get, expect := strings.TrimRight(rec.Body.String(), "\n"), fmt.Sprintf(fmtOutput, i); get != expect {
			t.Errorf("expect response body %q, get %q", expect, get)
		}
	}

}
