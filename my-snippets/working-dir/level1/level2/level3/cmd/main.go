package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"my-snippets/working-dir/config"
)

func main() {

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working directory: %v", err)
	}
	fmt.Printf("working dir path: %s\n", wd)

	ex, err := os.Executable()
	if err != nil {
		log.Fatalf("failed to get executable: %v", err)
	}
	fmt.Printf("executable path: %s\n", filepath.Dir(ex))

	fmt.Printf("config path: %s\n", config.ConfigPath)

}
