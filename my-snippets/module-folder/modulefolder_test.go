package modulefolder

import (
	"os"
	"path/filepath"
	"runtime"

	"testing"
)

func TestModuleFolder(t *testing.T) {

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("faield to get caller file name")
	}
	rootModuleDir := filepath.Dir(filename)

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current directory of test file: %v", err)
	}

	nestedFolderName := "nested0"

	err = os.Mkdir(nestedFolderName, 0755)
	if err != nil {
		t.Fatalf("failed to create nested folder: %v", err)
	}

	err = os.Chdir(filepath.Join(cwd, nestedFolderName))
	if err != nil {
		t.Fatalf("failed to change to nested folder: %v", err)
	}

	topModuleFolder, err := GetModuleFolder()
	if err != nil {
		t.Fatalf("failed to get module folder name from nested module folder: %v", err)
	}

	if topModuleFolder != rootModuleDir {
		t.Errorf("module folder:\nget\n%s\nexpect\n%s\n", topModuleFolder, rootModuleDir)
	}

	// cleanup
	os.Chdir(cwd)
	os.RemoveAll(nestedFolderName)
	os.Remove(nestedFolderName)

}
