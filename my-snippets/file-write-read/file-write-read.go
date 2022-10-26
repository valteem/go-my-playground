package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	txt := []byte("Some text to write to file")
	filename := "a.txt"
	e := os.WriteFile(filename, txt, 0600)
	if e != nil {
		fmt.Println("Error writing to file")
		return
	}

	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading from file")
		return
	}
	fmt.Println(string(input[:]))

	f, err := os.OpenFile("write.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	// 0644:
	// 6 - owner can read and write
	// 4 - users of the same group as user can read
	// 4 - all users can read
	// https://stackoverflow.com/questions/18415904/what-does-mode-t-0644-mean
	 
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err := f.Write([]byte("... logging some data ...\n")); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	
}