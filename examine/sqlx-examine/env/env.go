package env

import (
	"log"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

func init() {

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current directory from inside env.init() call: %v", err)
	}
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("failed to get caller filename")
	}
	err = os.Chdir(path.Dir(filename))
	if err != nil {
		log.Fatalf("failed to change directory: %v", err)
	}

	err = godotenv.Load()

	if err != nil {
		cwd, _ := os.Getwd()
		log.Fatalf("failed to load env variable(s), current directory is %s", cwd)
	}

	err = os.Chdir(cwd)
	if err != nil {
		log.Fatalf("failed to change directory back to original (%s) from inside env.init() call: %v", cwd, err)
	}

}
