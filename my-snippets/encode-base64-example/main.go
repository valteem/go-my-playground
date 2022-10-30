package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	msg := "tt"
	b := []byte(msg)
	for _, s := range(b) {
		fmt.Printf(" %06b", s)
	}
	fmt.Printf("\n")

	encoded := base64.StdEncoding.EncodeToString(b)
	fmt.Println(b, encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode error", err)
		return
	}
	fmt.Println(decoded, string(decoded))

}