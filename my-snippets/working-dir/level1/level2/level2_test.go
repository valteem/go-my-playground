package level2

import (
	"fmt"
	"os"
	"path/filepath"

	"testing"
)

func TestLevel2(t *testing.T) {

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("level 2: failed to get working directory: %v", err)
	}
	fmt.Printf("level 2: %s\n", wd)

	ex, err := os.Executable()
	if err != nil {
		t.Fatalf("level 2: failed to get executable: %v", err)
	}
	fmt.Printf("level 2 executable path: %s\n", filepath.Dir(ex))

}
