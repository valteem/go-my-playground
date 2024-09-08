package mware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func TestJWTHandler(t *testing.T) {

	keyFunc := func(context.Context) (any, error) {
		return []byte("secret"), nil
	}

	jwtValidator, err := validator.New(
		keyFunc,
		validator.HS256,
		"https://<issuer-url>/",
		[]string{"<audience>"},
	)
	if err != nil {
		t.Fatalf("failed to set up JWT validator: %v", err)
	}

	mware := jwtmiddleware.New(jwtValidator.ValidateToken)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJpc3MiOiJnby1qd3QtbWlkZGxld2FyZS1leGFtcGxlIiwiYXVkIjoiZ28tand0LW1pZGRsZXdhcmUtZXhhbXBsZSJ9.xcnkyPYu_b3qm2yeYuEgr5R5M5t4pN9s04U1ya53-KM")

	resp := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.Handle("/", mware.CheckJWT(jwtHandler))

	mux.ServeHTTP(resp, req)

	output := resp.Body.String()
	if len(output) == 0 {
		t.Errorf("get empty output")
	}
}
