package main

import (
	"log"
	"net/http"
)

func mwareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println("Running middleware #1 before next middleware")
		next.ServeHTTP(w, r)
		log.Println("Running middleware #1 after next middleware")
	})
}

func mwareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println("Running middleware #2 before next middleware")
		if r.URL.Path == "/skip" {
			return
		}
		next.ServeHTTP(w, r)
		log.Println("Running middleware #2 after next middleware")
	})
}

func appHandlerFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("Running app handler")
	w.Write([]byte("App Handler"))
}

func main() {

	mux := http.NewServeMux()

	appHandler := http.HandlerFunc(appHandlerFunc)

	mux.Handle("/", mwareOne(mwareTwo(appHandler)))

	log.Println("Listening ...")

	http.ListenAndServe(":3000", mux)
}