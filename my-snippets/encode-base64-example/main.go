package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	msg := "t0-a"

	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode error", err)
		return
	}
	fmt.Println(decoded, string(decoded))

}