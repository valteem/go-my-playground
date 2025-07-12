package onions

import (
	"os"
	"testing"

	paths "github.com/valteem/reuse/package-test-executable"
)

func TestOnions(t *testing.T) {

	fsegs, err := paths.ExecutablePathSegments()
	if err != nil {
		t.Fatalf("failed to get file path segments: %v", err)

	}
	topDirName, executableName := fsegs[1], fsegs[len(fsegs)-1]
	if topDirName != "tmp" || executableName != "apples.test" {
		t.Errorf("top directory and executable name:\n: get\n %s/%s\nexpect\n%s/%s",
			topDirName,
			executableName,
			"tmp",
			"apples.test",
		)
	}

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}
	fileContent, err := paths.GetAssets(cwd)
	if err != nil {
		t.Fatalf("failed to read file content: %v", err)
	}
	if fileContent != "some data" {
		t.Errorf("file content: get %q, expect %q", fileContent, "some data")
	}

}
