// go test -v from anchor/ directory
package anchor

import (
	"os"
	"strings"

	"testing"
)

func TestChangeDir(t *testing.T) {

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get path to current (anchor) dir: %v", err)
	}

	err = os.Chdir(strings.TrimSuffix(cwd, "anchor") + "next")
	if err != nil {
		t.Fatalf("failed to change folder: %v", err)
	}

	b, err := os.ReadFile("config.json")
	if err != nil {
		t.Fatalf("failed to read config: %v", err)
	}

	if len(b) == 0 {
		t.Errorf("config is empty")
	}
}

// os.Chdir() effect stays if both test are run in a single batch (go test)
// This test fails if it is run separately from the preious above
// This is why ChangeWorkingDir() from gotest.tools/v3 resets (by default) current directory
// under the hood after test functions it is invoked from is completed
func TestNextDir(t *testing.T) {

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get path to current (next?) dir: %v", err)
	}

	pathSegments := strings.Split(cwd, "/")
	folderName := pathSegments[len(pathSegments)-1]
	if folderName != "next" {
		t.Errorf("current folder name: get %q, expect %q", folderName, "next")
	}

}
