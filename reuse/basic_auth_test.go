package reuse_test

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasicAuth(t *testing.T) {

	var username, password string

	authHandler := func(w http.ResponseWriter, r *http.Request) {
		var ok bool
		username, password, ok = r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

	mux := http.DefaultServeMux
	mux.Handle("/user", http.HandlerFunc(authHandler))

	inputUserName, inputPassword := "some_user", "some_password"

	req := httptest.NewRequest(http.MethodPost, "/user", nil)
	resp := httptest.NewRecorder()

	auth := base64.StdEncoding.EncodeToString([]byte(inputUserName + ":" + inputPassword))
	req.Header.Add("Authorization", "Basic "+auth)

	mux.ServeHTTP(resp, req)

	if username != inputUserName || password != inputPassword {
		t.Errorf("username/password: get %s/%s, expect %s/%s", username, password, inputUserName, inputPassword)
	}

	if respCode := resp.Code; respCode != http.StatusOK {
		t.Errorf("response code: get %d, expect %d", respCode, http.StatusOK)
	}

}
