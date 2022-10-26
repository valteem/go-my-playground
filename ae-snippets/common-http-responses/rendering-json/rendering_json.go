package main

import (
	"encoding/json"
	"net/http"
)

type ThingFeatures struct {
	Thing string
	Features []string
}

func main() {

	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {

	thingfeatures := ThingFeatures{"Bottle", []string{"drink", "throw away"}}

	js, err := json.Marshal(thingfeatures)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500 Internal Server Error
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}