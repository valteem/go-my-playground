package env

import (
	"errors"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func init() {

	_, fpath, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("cannot fetch env file path"))
	}
	dir := filepath.Dir(fpath)
	envPath := filepath.Join(dir, ".env")

	if err := godotenv.Load(envPath); err != nil {
		panic(err)
	}
}
