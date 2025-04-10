package main

import (
	"testing"
	"time"
)

func TestSetAndGetRoles(t *testing.T) {

	rolesInput := "role1,role2,role3"

	tokenString, err := NewTokenWithRoles(rolesInput, time.Hour*24)
	if err != nil {
		t.Fatalf("failed to create token string: %v", err)
	}

	rolesOutput, err := GetRolesFromToken(tokenString)
	if err != nil {
		t.Fatalf("failed to get roles from token: %v", err)
	}
	if rolesOutput != rolesInput {
		t.Errorf("expect %q, get %q", rolesInput, rolesOutput)
	}

}
