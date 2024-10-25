package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func handleNumbers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	numbers, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("numbers: " + numbers))
}

func handleLetters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	letters, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("letters: " + letters))
}

func handleNumbersAndLetters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	values, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("numbers and letters: " + values))
}

func handleNoNumbersOrLetters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	values, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("something strange: " + values))
}
