package main

import "net/http"

func main() {

	http.HandleFunc("/", mux)
	http.ListenAndServe(":3000", nil)

}

func mux(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/accept":
		w.Header().Set("Server", "A Go Web Server - accept")
		w.WriteHeader(202)
	case "/success":
		w.Header().Set("Server", "A Go Web Server")
		w.WriteHeader(http.StatusOK) // 200 OK
	case "/bullshit":
		w.Header().Set("Bullshit", "A lot of bullshit")
		w.WriteHeader(901)
	}

}