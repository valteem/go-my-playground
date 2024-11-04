package sessions

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
)

var (
	storeAuthKey = []byte("authentication_key") // for testing purposes only

	// The encryption key, if set, must be either 16, 24, or 32 bytes
	// storeEncryptKey = []byte("encryption_key")
	// store           = sessions.NewCookieStore(storeAuthKey, storeEncryptKey)

	store = sessions.NewCookieStore(storeAuthKey)
)

func testHandler(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "my-session")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if sessionNameExpected := "my-session"; session.Name() != sessionNameExpected {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	session.Values["user"] = "some user"
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TestNewSession(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/login", nil)
	resp := httptest.NewRecorder()
	mux := http.NewServeMux()
	mux.Handle("/login", http.HandlerFunc(testHandler))

	mux.ServeHTTP(resp, req)

	c := resp.Result().Cookies()
	if len(c) == 0 {
		t.Fatalf("failed to fetch any cookies")
	}

}
