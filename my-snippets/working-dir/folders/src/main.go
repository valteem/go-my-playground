package main

import (
	"log"
	"os"
)

const (
	fileContentExpected = "some content"
)

func main() {

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to ger current directory: %v", err)
	}

	log.Printf("current directory: %s", cwd)

	b, err := os.ReadFile("config.txt")
	if err != nil {
		log.Fatalf("failed to read file content: %v", err)
	}

	if fileContentActual := string(b); fileContentActual != fileContentExpected {
		log.Fatalf("file content: get %q, expect %q", fileContentActual, fileContentExpected)
	}

}
