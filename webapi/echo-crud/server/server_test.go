package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

var (
	createJSON       = `{"desc":"someNewItem"}`
	updateJSON       = `{"desc":"newItemDescription"}`
	fmtOutputCreated = `{"id":%d,"desc":"someNewItem"}`
	fmtOutputUpdated = `{"id":%d,"desc":"newItemDescription"}`
)

func TestCreateItem(t *testing.T) {

	e := echo.New()

	for i := 1; i < 11; i++ {
		req := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(createJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) // Content-Type: application/json
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		if err := CreateItem(c); err != nil {
			t.Fatalf("error creating item: %v", err)
		}
		if rec.Code != http.StatusCreated {
			t.Errorf("expect response code %d, get %d", http.StatusCreated, rec.Code)
		}
		// https://github.com/labstack/echo/discussions/2450
		if get, expect := strings.TrimRight(rec.Body.String(), "\n"), fmt.Sprintf(fmtOutputCreated, i); get != expect {
			t.Errorf("expect response body %q, get %q", expect, get)
		}
	}

	for i := 1; i < 15; i++ {
		target := "/items/" + strconv.Itoa(i)
		req := httptest.NewRequest(http.MethodGet, target, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/items/:id") // test routing
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(i))
		if err := GetItem(c); err != nil {
			t.Fatalf("error getting item: %v", err)
		}
		var expectedStatus int
		var expectedBody string
		if i < 11 { // existing items
			expectedStatus = http.StatusOK
			expectedBody = fmt.Sprintf(fmtOutputCreated, i)
		} else { // non-existing items
			expectedStatus = http.StatusNotFound
			expectedBody = "null"
		}
		if rec.Code != expectedStatus {
			t.Errorf("expect response code %d, get %d", expectedStatus, rec.Code)
		}
		getBody := strings.TrimRight(rec.Body.String(), "\n")
		if getBody != expectedBody {
			t.Errorf("expect response body %q, get %q", expectedBody, getBody)
		}
	}

	for i := 6; i < 15; i++ {
		target := "/items/" + strconv.Itoa(i)
		req := httptest.NewRequest(http.MethodPut, target, strings.NewReader(updateJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) // Content-Type: application/json
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/items/:id") // test routing
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(i))
		if err := UpdateItem(c); err != nil {
			t.Fatalf("error updating item: %v", err)
		}
		var expectedStatus int
		var expectedBody string
		if i < 11 { // existing items
			expectedStatus = http.StatusOK
			expectedBody = fmt.Sprintf(fmtOutputUpdated, i)
		} else { // non-existing items
			expectedStatus = http.StatusNotFound
			expectedBody = "null"
		}
		if rec.Code != expectedStatus {
			t.Errorf("expect response code %d, get %d", expectedStatus, rec.Code)
		}
		getBody := strings.TrimRight(rec.Body.String(), "\n")
		if getBody != expectedBody {
			t.Errorf("expect response body %q, get %q", expectedBody, getBody)
		}
	}

	for i := 8; i < 12; i++ {
		target := "/items/" + strconv.Itoa(i)
		req := httptest.NewRequest(http.MethodDelete, target, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/items/:id") // test routing
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(i))
		if err := DeleteItem(c); err != nil {
			t.Fatalf("error updating item: %v", err)
		}
		var expectedStatus int
		var expectedBody string
		if i < 11 { // existing items
			expectedStatus = http.StatusNoContent
			expectedBody = ""
		} else { // non-existing items
			expectedStatus = http.StatusNotFound
			expectedBody = "null"
		}
		if rec.Code != expectedStatus {
			t.Errorf("expect response code %d, get %d", expectedStatus, rec.Code)
		}
		getBody := strings.TrimRight(rec.Body.String(), "\n")
		if getBody != expectedBody {
			t.Errorf("expect response body %q, get %q", expectedBody, getBody)
		}
	}

}
