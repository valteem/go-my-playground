package main

import (
	"encoding/json"
	"net/http"
)

func HandlePersonData(p *Person) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(p)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// http.ResponseWriter.Write() implicitly calls w.WriteHeader(http.StatusOK)
		//		w.WriteHeader(http.StatusOK)
	})
}
