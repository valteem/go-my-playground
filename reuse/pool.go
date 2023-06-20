package reuse

import (
	"fmt"
)

type Resource struct {
	Name string 
}

func New() any {
	fmt.Println("new")
	return &Resource{"resource"}
}