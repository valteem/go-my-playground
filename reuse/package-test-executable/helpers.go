package packagetestexecutable

import (
	"os"
	"path"
	"strings"
)

func ExecutablePathSegments() ([]string, error) {

	fpath, err := os.Executable()
	if err != nil {
		return make([]string, 0), err
	}

	return strings.Split(fpath, "/"), nil

}

func GetAssets(cwd string) (string, error) {

	pathToAssets := path.Join(cwd, "..", "testdata", "file.txt")

	b, err := os.ReadFile(pathToAssets)
	if err != nil {
		return "", err
	}

	return string(b), nil

}
