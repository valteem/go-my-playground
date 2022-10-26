package main

import (
	"fmt"
	"encoding/json"
)

type Country struct {
    Name string `json:"name"`
    Index int `json:"index"`
}

type Countries struct {
    Data []Country `json:"countries"`
}

func main() {

	data := []byte(`
		{"countries":
            [{"name":"Belgium","index":20}, 
             {"name":"France","index":17}]
        }`)


	var parsed Countries
	json.Unmarshal([]byte(data), &parsed)

	fmt.Println(parsed)

}
