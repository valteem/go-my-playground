package reuse_test

import (
	"os"
	"testing"
)

func TestSetEnv(t *testing.T) {
	envVarName := "SOME_ENV_VAR"
	envVarValue := "TRUE"
	err := os.Setenv(envVarName, envVarValue)
	if err != nil {
		t.Fatalf("failed to set env variable: %q", err)
	}
	actualEnvVarValue := os.Getenv(envVarName)
	if actualEnvVarValue != envVarValue {
		t.Errorf("os.getenv(%s): get %s, expect %s", envVarName, actualEnvVarValue, envVarValue)
	}
}
