package main

import (
	"encoding/json"
	"log"
)

type CrateWithToys struct {
	ID uint32		`json:"id"`
	Name string		`json:"name"`
	Toys []string	`json:"toys"`
}

func main() {
	var cwt []CrateWithToys // SLICE of struct instances !!!
	b := []byte(`[
	{"id": 1, "name": "New crate", "toys": ["Car", "Ball", "Pen"]},
	{"id": 2, "name": "Small crate", "toys": ["nothing"]}
	]`)

	err := json.Unmarshal(b, &cwt)
	if err != nil {
		log.Println("Error unmarshalling json blob", err)
		return
	}
	log.Println(cwt)
}