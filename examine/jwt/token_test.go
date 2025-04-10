package main

import (
	"reflect"
	"testing"
	"time"
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
