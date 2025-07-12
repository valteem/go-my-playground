package modulefolder

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetModuleFolder() (string, error) {

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to fetch caller filename")
	}

	currentDir := filepath.Dir(filename)

	for {

		goModPath := filepath.Join(currentDir, "go.mod")
		_, err := os.Stat(goModPath)
		if err == nil {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return "", fmt.Errorf("reached root folder, no go.mod found")
		}

		currentDir = parentDir
	}

}
