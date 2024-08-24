package echo

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestEchoBasic(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	if e.Router() == nil {
		t.Errorf("echo router cannot be nil")
	}

	c := e.NewContext(req, res)
	e.DefaultHTTPErrorHandler(errors.New("error"), c)
	if res.Code != http.StatusInternalServerError {
		t.Errorf("expect %d error code, get %d", http.StatusInternalServerError, res.Code)
	}

}
