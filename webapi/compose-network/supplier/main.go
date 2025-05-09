package main

import (
	"encoding/json"
	"hash/crc64"
	"net/http"
)

type responseMsg struct {
	CheckSum uint64 `json:"check_sum"`
}

func handleCheckSum(w http.ResponseWriter, r *http.Request) {

	input := r.PathValue("input")

	sum := crc64.Checksum([]byte(input), crc64.MakeTable(crc64.ISO))
	jsonMsg, err := json.Marshal(responseMsg{CheckSum: sum})
	if err != nil {
		w.Write([]byte("internal error"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(jsonMsg))
	w.WriteHeader(http.StatusOK)
}

func main() {

	http.Handle("/checksum/{input}", http.HandlerFunc(handleCheckSum))

	http.ListenAndServe(":3002", nil)

}
