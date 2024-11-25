package auth

import (
	"fmt"
	"net/http"
	// "webapi/wrap-auth/handlers"
)

const (
	accessGranted = "authorized"
)

var (
	ErrAuthUser = fmt.Errorf("user not authorized")
)

func AuthorizeUser(token string) error {
	if token == accessGranted {
		return nil
	}
	return ErrAuthUser
}

func GrantAccess(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		if err := AuthorizeUser(authToken); err != nil {
			http.Error(w, "auth token not valid", http.StatusForbidden)
			return
		}
		fn(w, r)
	}
}
