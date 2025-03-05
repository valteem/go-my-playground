package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var (
	vars = []struct {
		varname  string
		varvalue string
	}{
		{"SOME_ENV_VAR", "VALUE1"},
		{"SOME_OTHER_ENV_VAR", "VALUE2"},
		{"SOME_RANDOM_VAR", "VALUE3"},
	}
)

func TestEnvVar(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		t.Fatalf("failed to load env variables: %v", err)
	}

	for _, v := range vars {
		if value := os.Getenv(v.varname); value != v.varvalue {
			t.Errorf("%q value: get %q, expect %q", v.varname, value, v.varvalue)
		}
	}

}
