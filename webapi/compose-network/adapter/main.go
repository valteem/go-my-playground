package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	client            *http.Client
	supplierURLPrefix string
)

func init() {

	client = &http.Client{}

	supplierURLPrefix = os.Getenv("SUPPLIER_URL_PREFIX")

}

func handleInput(w http.ResponseWriter, r *http.Request) {

	input := r.PathValue("input")

	req, err := http.NewRequest(http.MethodGet, "http://"+supplierURLPrefix+"/checksum/"+input, nil)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("internal error: %v", err)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("internal error: %v", err)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("internal error: %v", err)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(respBody)
	w.WriteHeader(http.StatusOK)

}

func main() {

	http.Handle("/input/{input}", http.HandlerFunc(handleInput))

	http.ListenAndServe(":3001", nil)

}
