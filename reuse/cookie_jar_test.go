package reuse_test

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "some-name",
		Value: "some_value",
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(200)
}

func readCookie(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("some-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(cookie.Name + " " + cookie.Value))
	w.WriteHeader(http.StatusOK)

}

func runServer() {

	mux := http.NewServeMux()

	mux.Handle("/set", http.HandlerFunc(setCookie))
	mux.Handle("/read", http.HandlerFunc(readCookie))

	http.ListenAndServe(":3001", mux)

}

func TestCookieJar(t *testing.T) {

	jar, err := cookiejar.New(nil) //no options
	if err != nil {
		t.Fatalf("failed to create new cookie jar: %v", err)
	}

	client := &http.Client{
		Jar: jar,
	}

	go runServer()

	time.Sleep(1 * time.Second)

	respSet, err := client.Get("http://127.0.0.1:3001/set")
	if err != nil {
		t.Fatalf("failed to get response (/set): %v", err)
	}
	defer respSet.Body.Close()

	cookies := respSet.Cookies()
	if len(cookies) == 0 {
		t.Fatalf("failed to fetch any cookies upon create")
	}
	respRead, err := client.Get("http://127.0.0.1:3001/read")
	if err != nil {
		t.Fatalf("failed to get response (/read): %v", err)
	}
	defer respRead.Body.Close()

	body, err := io.ReadAll(respRead.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	if actual, expected := string(body), "some-name some_value"; actual != expected {
		t.Errorf("response to \"/read\" body:\nget\n%q\nexpect\n%q", actual, expected)
	}

}
