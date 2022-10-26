package main

import "fmt"
import "example.com/example-auth-local/auth_local"

func main() {
	var k [32]byte
	a := []byte("authenticator")
	copy(k[:], []byte("secret-key"))
	fmt.Println(a)
	fmt.Println(k)
	s := auth_local.SumLocal(a, &k)
	fmt.Println(*s)
}