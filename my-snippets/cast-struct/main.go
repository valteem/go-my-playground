package main

import (
	"encoding/json"
	"fmt"
)

type ErrorMessage struct {
	Text string
}

type T struct {
	Text string `json:"msg"`
}

func main() {

	var e1 ErrorMessage
	var e2 ErrorMessage
	
	data := []byte(`{"msg": "some text"}`)

// Type ErrorMessage does not contain json annotations, so an object of this type cannot be assigned by json.Unmarshal
// This returns empty error message
	if json.Unmarshal(data, &e1)==nil {fmt.Println(e1)}
// This returns 'some text' error message, due to casting e2 as T, which has json annotation	
	if json.Unmarshal(data, (*T)(&e2))==nil {fmt.Println(e2)}
	
}