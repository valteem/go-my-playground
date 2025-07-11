package callerdir

import (
	"log"
	"os"
	"path"
	"runtime"
)

// https://stackoverflow.com/a/60258660
func init() {

	_, filename, _, _ := runtime.Caller(0)
	err := os.Chdir(path.Dir(filename))
	if err != nil {
		log.Fatalf("failed to change directory: %v", err)
	}

}
