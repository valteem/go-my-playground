package models

import "testing"

func TestUserConstants(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		actual   string
	}{
		{"RoleAdmin", "admin", RoleAdmin},
		{"RoleGreetingsManager", "greetings_manager", RoleGreetingsManager},
		{"RoleUser", "user", RoleUser},
		{"PermissionAddGreeting", "greeting:add", PermissionAddGreeting},
		{"PermissionManageUsers", "user:manage", PermissionManageUsers},
		{"PermissionGetGreeting", "greeting:get", PermissionGetGreeting},
		{"PermissionManageGreetings", "greeting:manage", PermissionManageGreetings},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.actual != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, tt.actual)
			}
		})
	}
}

func TestUserStruct(t *testing.T) {
	user := User{
		ID:       1,
		Username: "testuser",
		Password: "hashedpassword",
		Role:     RoleUser,
	}

	if user.ID != 1 {
		t.Errorf("Expected ID 1, got %d", user.ID)
	}

	if user.Username != "testuser" {
		t.Errorf("Expected Username 'testuser', got '%s'", user.Username)
	}

	if user.Role != RoleUser {
		t.Errorf("Expected Role '%s', got '%s'", RoleUser, user.Role)
	}
}
