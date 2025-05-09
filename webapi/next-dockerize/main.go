package main

import (
	"encoding/json"
	"net/http"
)

type responseMsg struct {
	Msg string `json:"response_message"`
}

var (
	msg []string
	idx int
)

func init() {
	msg = make([]string, 0)
	msg = append(msg, "some random message")
	msg = append(msg, "another random message")
	msg = append(msg, "yet another random message")
	idx = 0
}

func handleNext(w http.ResponseWriter, r *http.Request) {
	m := responseMsg{Msg: msg[idx]}
	idx++
	if idx == len(msg) {
		idx = 0
	}
	jsonMsg, err := json.Marshal(m)
	if err != nil {
		w.Write([]byte("internal error"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(jsonMsg))
	w.WriteHeader(http.StatusOK)
}

func main() {

	http.Handle("/next", http.HandlerFunc(handleNext))

	http.ListenAndServe(":3001", nil)

}
