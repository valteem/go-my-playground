package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	PartNum string `json:"partnum"`
	Store   string `json:"store"`
	Qty     int    `json:"qty"`
}

// don't need separate handlers for GET and POST, since r.ParseForm does the trick for both
/*
func handleGet(w http.ResponseWriter, r *http.Request) {
	var response Response
	response.PartNum = r.URL.Query().Get("partnum")
	response.Store = r.URL.Query().Get("store")

	sendResponse(w, &response)
}
*/

func handleBothGetAndPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() // fills in both r.Form and r.PostForm
	if err != nil {
		log.Printf("Error parsing request data: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	partNum := r.Form["partnum"][0]
	store := r.Form["store"][0]
	response := Response{PartNum: partNum, Store: store}

	sendResponse(w, &response)
}

func sendResponse(w http.ResponseWriter, response *Response) {

	getQty(response)

	data, err := json.Marshal(&response)
	if err != nil {
		log.Printf("Error marshalling response data to JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func getQty(response *Response) {
	response.Qty = 1
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/api/query", handleBothGetAndPost)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
