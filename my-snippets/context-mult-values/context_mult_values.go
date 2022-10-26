package main

import (
	"fmt"
	"context"
	"reflect"
)

func main() {

// Key must be comparable and should not be of type string or any other built-in type to avoid collisions 
// between packages using context. Users of WithValue should define their own types for keys
	type contextKey string
	ctxKey := contextKey("intnum")

	ctxOne := context.WithValue(context.Background(), ctxKey, "one")
	ctxTwo := context.WithValue(ctxOne, ctxKey, "two")
	ctxThree := context.WithValue(ctxTwo, ctxKey, "three")
	fmt.Println(ctxOne.Value(ctxKey))
	fmt.Println(ctxTwo.Value(ctxKey))
	fmt.Println(ctxThree.Value(ctxKey))

// Some obscure experiments with context key and value types
	keys := reflect.TypeOf(ctxThree).Elem()
	fmt.Println(keys)
	values := reflect.ValueOf(ctxThree).Elem()
	fmt.Println(values)
}