// https://stackoverflow.com/q/73947665
package reuse_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type ctxKey uint8

const userKey ctxKey = 0

type user struct{ name string }

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("log")

		u := new(user)                                                // avoid pointer-to-a-string dereferencing
		r = r.WithContext(context.WithValue(r.Context(), userKey, u)) // change incoming request

		defer func(start time.Time) {
			if u.name != "" {
				fmt.Printf("user %s has accessed %s, took %s\n", u.name, r.URL.Path, time.Since(start))
			} else {
				fmt.Printf("anonimous has accessed %s, took %s\n", r.URL.Path, time.Since(start))
			}
		}(time.Now())

		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if u, ok := r.Context().Value(userKey).(*user); ok {
			u.name = "user123"
		}
		fmt.Println("auth")
		next.ServeHTTP(w, r)
	})
}

func welcome(w http.ResponseWriter, r *http.Request) {
	time.Sleep(4567) // arbitrary sleep to see that logMiddleware does its job as expected

	if u, ok := r.Context().Value(userKey).(*user); ok && u.name != "" {
		fmt.Fprintf(w, "hello %s", u.name)
	} else {
		fmt.Fprintf(w, "hello")
	}
}

func TestContextChain(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/welcome", welcome)
	chain := logMiddleware(authMiddleware(mux))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/welcome", nil)
	chain.ServeHTTP(w, r)

	//	w.Flush()
	if bodyActual, bodyExpected := w.Body.String(), "hello user123"; bodyActual != bodyExpected {
		t.Errorf("response body: get %s, expect %s", bodyActual, bodyExpected)
	}
}
