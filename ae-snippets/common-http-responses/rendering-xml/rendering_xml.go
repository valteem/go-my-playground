package main

import (
	"encoding/xml"
	"net/http"
)

type ThingFeatures struct{
	Thing string
	Features []string `xml:"Features>Feature"`
}

func main() {

	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {

	thingfeature := ThingFeatures{"Bottle", []string{"drink", "throw away"}}

	x, err := xml.MarshalIndent(thingfeature, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500 Internal Server Error
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)

}