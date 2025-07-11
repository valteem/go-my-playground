package entry

import (
	"log"
	"os"

	"testing"
)

func TestConfigDirAccess(t *testing.T) {

	// Relative path works for running 'go test' from this and any parent folder
	b, err := os.ReadFile("../config/config.json")
	if err != nil {
		t.Fatalf("failed to read config file: %v", err)
	}

	fpath, _ := os.Executable()
	log.Println(fpath)

	if len(b) == 0 {
		t.Errorf("config file empty")
	}

}
