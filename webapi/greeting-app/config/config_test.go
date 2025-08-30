package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Test default values
	cfg := LoadConfig()

	if cfg.DBHost == "" {
		t.Error("DBHost should not be empty")
	}
	if cfg.DBPort == "" {
		t.Error("DBPort should not be empty")
	}
	if cfg.DBUser == "" {
		t.Error("DBUser should not be empty")
	}
	if cfg.DBName == "" {
		t.Error("DBName should not be empty")
	}
	if cfg.JWTSecret == "" {
		t.Error("JWTSecret should not be empty")
	}
}

func TestGetEnv(t *testing.T) {
	// Test with existing environment variable
	os.Setenv("TEST_VAR", "test_value")
	defer os.Unsetenv("TEST_VAR")

	value := getEnv("TEST_VAR", "default")
	if value != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", value)
	}

	// Test with non-existing environment variable
	value = getEnv("NON_EXISTING_VAR", "default_value")
	if value != "default_value" {
		t.Errorf("Expected 'default_value', got '%s'", value)
	}
}
