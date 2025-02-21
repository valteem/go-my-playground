package main

import (
	"compress/gzip"
	"encoding/json"
	"net/http"
)

func HandleGZipResponse(s string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		body, err := json.Marshal(s)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer, err := gzip.NewWriterLevel(w, gzip.DefaultCompression)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer writer.Close()

		writer.Write(body)

	}

}
