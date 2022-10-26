package main

import (
	"encoding/json"
	"log"
	"os"
)

type BookShelf struct {
	ID uint32        `json:"id"`
	Name string      `json:"name"`
	Books []string   `json:"books"`
}

func main() {
	b := BookShelf{
		ID: 501,
		Name: "First shelf",
		Books: []string{"Good book", "Average book", "Old book", "Strange book", "Yet another book"},
	}

	m, _ := json.Marshal(b)
	err := os.WriteFile("Books.json", []byte(m), 0666)
	if err != nil {
		log.Println("Error writing json file")
	}
}