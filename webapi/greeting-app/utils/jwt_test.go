package utils

import (
	"greeting-app/config"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateJWT(t *testing.T) {
	cfg := &config.Config{JWTSecret: "test_secret"}

	token, err := GenerateJWT(1, "testuser", cfg)
	if err != nil {
		t.Fatalf("GenerateJWT failed: %v", err)
	}

	if token == "" {
		t.Error("Generated token should not be empty")
	}
}

func TestValidateJWT(t *testing.T) {
	cfg := &config.Config{JWTSecret: "test_secret"}

	// Generate a valid token
	token, err := GenerateJWT(1, "testuser", cfg)
	if err != nil {
		t.Fatalf("GenerateJWT failed: %v", err)
	}

	// Validate the token
	claims, err := ValidateJWT(token, cfg)
	if err != nil {
		t.Fatalf("ValidateJWT failed: %v", err)
	}

	if claims.UserID != 1 {
		t.Errorf("Expected UserID 1, got %d", claims.UserID)
	}

	if claims.Username != "testuser" {
		t.Errorf("Expected Username 'testuser', got '%s'", claims.Username)
	}
}

func TestValidateJWTInvalidToken(t *testing.T) {
	cfg := &config.Config{JWTSecret: "test_secret"}

	// Test invalid token
	_, err := ValidateJWT("invalid.token.string", cfg)
	if err == nil {
		t.Error("ValidateJWT should fail with invalid token")
	}

	// Test token with wrong secret
	validCfg := &config.Config{JWTSecret: "correct_secret"}
	wrongCfg := &config.Config{JWTSecret: "wrong_secret"}

	token, err := GenerateJWT(1, "testuser", validCfg)
	if err != nil {
		t.Fatalf("GenerateJWT failed: %v", err)
	}

	_, err = ValidateJWT(token, wrongCfg)
	if err == nil {
		t.Error("ValidateJWT should fail with wrong secret")
	}
}

func TestJWTExpiration(t *testing.T) {
	cfg := &config.Config{JWTSecret: "test_secret"}

	// Create a token that expires in 1 second
	expirationTime := time.Now().Add(1 * time.Second)
	claims := &Claims{
		UserID:   1,
		Username: "testuser",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		t.Fatalf("Failed to create token: %v", err)
	}

	// Token should be valid immediately
	_, err = ValidateJWT(tokenString, cfg)
	if err != nil {
		t.Errorf("Token should be valid immediately: %v", err)
	}

	// Wait for token to expire
	time.Sleep(2 * time.Second)

	// Token should now be expired
	_, err = ValidateJWT(tokenString, cfg)
	if err == nil {
		t.Error("Token should be expired after waiting")
	}

	// Check that it's specifically an expiration error
	if err != nil && err != jwt.ErrTokenExpired {
		t.Logf("Expected expiration error, got: %v", err)
	}
}

func TestJWTWithPastExpiration(t *testing.T) {
	cfg := &config.Config{JWTSecret: "test_secret"}

	// Create a token that expired 1 second ago
	expirationTime := time.Now().Add(-1 * time.Second)
	claims := &Claims{
		UserID:   1,
		Username: "testuser",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		t.Fatalf("Failed to create token: %v", err)
	}

	// Token should be expired
	_, err = ValidateJWT(tokenString, cfg)
	if err == nil {
		t.Error("Token with past expiration should fail validation")
	}

	if err == jwt.ErrTokenExpired {
		t.Log("Correctly identified expired token")
	}
}
