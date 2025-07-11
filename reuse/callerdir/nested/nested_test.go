package nested

import (
	"os"

	"testing"

	_ "github.com/valteem/reuse/callerdir"
)

func TestReadConfig(t *testing.T) {

	cwd, _ := os.Getwd()
	t.Log(cwd)

	b, err := os.ReadFile("config/config.json") // path from reuse/ folder
	if err != nil {
		t.Fatalf("failed to read config file from root project dir: %v", err)
	}

	if len(b) == 0 {
		t.Errorf("get empty config")
	}

}
