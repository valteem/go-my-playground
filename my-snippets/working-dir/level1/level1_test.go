package level1

import (
	"fmt"
	"os"
	"path/filepath"

	"testing"
)

func TestLevel1(t *testing.T) {

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("level 1: failed to get working directory: %v", err)
	}
	fmt.Printf("level 1 test path: %s\n", wd)

	ex, err := os.Executable()
	if err != nil {
		t.Fatalf("level 1: failed to get executable: %v", err)
	}
	fmt.Printf("level 1 executable path: %s\n", filepath.Dir(ex))

}
