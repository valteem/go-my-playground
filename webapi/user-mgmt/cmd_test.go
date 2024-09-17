package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {

	initStorage()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/users/register", registerUser)
	mux.HandleFunc("/api/users/authenticate", authenticateUser)

	// var bodyReader strings.Reader
	// _, err := bodyReader.Read([]byte("{\"login\":\"first_user\",\"password\":\"first_user_password\"}"))
	bodyReader := strings.NewReader("{\"login\":\"first_user\",\"password\":\"first_user_password\"}")
	/* 	if err != nil {
	   		t.Fatalf("error reading request body string: %v", err)
	   	}
	*/
	req := httptest.NewRequest(http.MethodPost, "/api/users/register", bodyReader)
	resp := httptest.NewRecorder()

	mux.ServeHTTP(resp, req)

	if _, ok := credentials["first_user"]; !ok {
		t.Errorf("failed to add new user")
	}

}
