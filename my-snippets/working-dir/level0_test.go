package main

import (
	"fmt"
	"os"
	"path/filepath"

	"testing"
)

func TestLevel0(t *testing.T) {

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("level 0: failed to get working directory: %v", err)
	}
	fmt.Printf("level 0: %s\n", wd)

	ex, err := os.Executable()
	if err != nil {
		t.Fatalf("level 0: failed to get executable: %v", err)
	}
	fmt.Printf("level 0 executable path: %s\n", filepath.Dir(ex))

}
