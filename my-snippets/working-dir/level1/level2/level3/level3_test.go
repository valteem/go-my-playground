package level3

import (
	"fmt"
	"os"
	"path/filepath"

	"testing"
)

func TestLevel3(t *testing.T) {

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("level 3: failed to get working directory: %v", err)
	}
	fmt.Printf("level 3: %s\n", wd)

	ex, err := os.Executable()
	if err != nil {
		t.Fatalf("level 3: failed to get executable: %v", err)
	}
	fmt.Printf("level 2 executable path: %s\n", filepath.Dir(ex))

}
