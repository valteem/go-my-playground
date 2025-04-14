package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestSetAndGetRoles(t *testing.T) {

	input := map[string]string{
		"claim1": "value1",
		"claim2": "value2",
		"claim3": "value3",
	}

	tokenString, err := NewTokenWithCustomClaims(input, time.Hour*24)
	if err != nil {
		t.Fatalf("failed to create token string: %v", err)
	}

	output, err := GetCustomClaimsFromToken(tokenString, []string{
		"claim1",
		"claim2",
		"claim3",
	})
	if err != nil {
		t.Fatalf("failed to get roles from token: %v", err)
	}
	if !reflect.DeepEqual(output, input) {
		t.Errorf("expect %v, get %v", input, output)
	}

}

func TestTokenClaims(t *testing.T) {

	userId := "universe42"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(defaultTokenTTL)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
		UserId: userId,
	})

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		t.Fatalf("failed to encrypt token: %v", err)
	}

	outputToken, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidTokenSigningMethod
		}
		return privateKey, nil
	})
	if err != nil {
		t.Fatalf("failed to decrypt the signed token: %v", err)
	}

	outputClaims, ok := outputToken.Claims.(*TokenClaims)
	if !ok {
		t.Fatalf("failed to bring output token claims to TokenClaims")
	}
	if outputClaims.UserId != userId {
		t.Errorf("UserId: get %q, expect %q", outputClaims.UserId, userId)
	}

}
