package rbac

import (
	"testing"

	"greeting-app/models"
)

func TestHasPermission(t *testing.T) {
	tests := []struct {
		role       string
		permission string
		expected   bool
		name       string
	}{
		{models.RoleAdmin, models.PermissionAddGreeting, true, "Admin should have add greeting permission"},
		{models.RoleAdmin, models.PermissionManageUsers, true, "Admin should have manage users permission"},
		{models.RoleGreetingsManager, models.PermissionAddGreeting, true, "Greetings manager should have add greeting permission"},
		{models.RoleGreetingsManager, models.PermissionManageUsers, false, "Greetings manager should NOT have manage users permission"},
		{models.RoleUser, models.PermissionGetGreeting, true, "User should have get greeting permission"},
		{models.RoleUser, models.PermissionAddGreeting, false, "User should NOT have add greeting permission"},
		{models.RoleUser, "nonexistent:permission", false, "Non-existent permission should return false"},
		{"nonexistent_role", models.PermissionGetGreeting, false, "Non-existent role should return false"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HasPermission(tt.role, tt.permission)
			if result != tt.expected {
				t.Errorf("HasPermission(%s, %s) = %v, want %v", tt.role, tt.permission, result, tt.expected)
			}
		})
	}
}

func TestGetRolePermissions(t *testing.T) {
	// Test admin role
	adminPerms := GetRolePermissions(models.RoleAdmin)
	if len(adminPerms) == 0 {
		t.Error("Admin should have permissions")
	}

	// Test user role
	userPerms := GetRolePermissions(models.RoleUser)
	if len(userPerms) == 0 {
		t.Error("User should have permissions")
	}

	// Test non-existent role
	nonExistentPerms := GetRolePermissions("nonexistent_role")
	if len(nonExistentPerms) != 0 {
		t.Error("Non-existent role should have no permissions")
	}
}

func TestGetAllRolePermissions(t *testing.T) {
	allRolePerms := GetAllRolePermissions()

	// Should have at least our three roles
	if len(allRolePerms) < 3 {
		t.Errorf("Expected at least 3 roles, got %d", len(allRolePerms))
	}

	// Check that each role has proper structure
	for _, rolePerm := range allRolePerms {
		if rolePerm.Role == "" {
			t.Error("Role should not be empty")
		}
		if len(rolePerm.Permissions) == 0 {
			t.Errorf("Role %s should have permissions", rolePerm.Role)
		}
	}
}
