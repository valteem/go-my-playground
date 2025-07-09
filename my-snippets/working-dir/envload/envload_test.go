package envload

import (
	"os"
	"strings"

	"testing"

	"github.com/joho/godotenv"
)

const (
	envVarName  = "SOME_VAR"
	envVarValue = "SOME_VALUE"
)

func TestRelativePath(t *testing.T) {

	if err := godotenv.Load("../env/.env"); err != nil {
		t.Fatalf("failed to load from sibling folder using relative path: %v", err)
	}

	if value := os.Getenv(envVarName); value != envVarValue {
		t.Errorf("%q value^ get %q, expect %q", envVarName, value, envVarValue)
	}

}

func TestAbsolutePath(t *testing.T) {

	path, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get path to current folder: %v", err)
	}

	path = strings.TrimSuffix(path, "envload") + "env/.env"

	if err := godotenv.Load(path); err != nil {
		t.Fatalf("failed to load from sibling folder using absolute path: %v", err)
	}

	if value := os.Getenv(envVarName); value != envVarValue {
		t.Errorf("%q value^ get %q, expect %q", envVarName, value, envVarValue)
	}

}
