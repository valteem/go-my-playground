package main

import (
	"net/http"
)

func servePage(page []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(page)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "text/html")
	}
}

func main() {

	p := `
	<h1>Top header<h1>
	<h2>2nd level header<h2>
	<a>Just some plain text here<a>
	`

	http.Handle("/page", http.HandlerFunc(servePage([]byte(p))))

	http.ListenAndServe(":3001", nil)

}
