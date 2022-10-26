package main

import(
	"fmt"
)

func main() {

	message := make(chan string, 2)

	message <- "buff1"
	message <- "buff2"

	fmt.Println(<-message)
	fmt.Println(<-message)
}

