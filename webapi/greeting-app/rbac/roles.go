package rbac

import "greeting-app/models"

// Role-based permissions mapping
var RolePermissions = map[string][]string{
	models.RoleAdmin: {
		models.PermissionAddGreeting,
		models.PermissionManageUsers,
		models.PermissionGetGreeting,
		models.PermissionManageGreetings,
	},
	models.RoleGreetingsManager: {
		models.PermissionAddGreeting,
		models.PermissionGetGreeting,
		models.PermissionManageGreetings,
	},
	models.RoleUser: {
		models.PermissionGetGreeting,
	},
}

// Permission descriptions
var PermissionDescriptions = map[string]string{
	models.PermissionAddGreeting:     "Add new greetings to the system",
	models.PermissionManageUsers:     "Manage user accounts and roles",
	models.PermissionGetGreeting:     "Get random greetings",
	models.PermissionManageGreetings: "Manage existing greetings",
}

// Check if a role has a specific permission
func HasPermission(role, permission string) bool {
	permissions, exists := RolePermissions[role]
	if !exists {
		return false
	}

	for _, perm := range permissions {
		if perm == permission {
			return true
		}
	}
	return false
}

// Get all permissions for a role
func GetRolePermissions(role string) []models.Permission {
	var permissions []models.Permission
	permList, exists := RolePermissions[role]
	if !exists {
		return permissions
	}

	for _, perm := range permList {
		permissions = append(permissions, models.Permission{
			Name:        perm,
			Description: PermissionDescriptions[perm],
		})
	}
	return permissions
}

// Get all available roles and their permissions
func GetAllRolePermissions() []models.RolePermissions {
	var rolePerms []models.RolePermissions

	for role := range RolePermissions {
		rolePerms = append(rolePerms, models.RolePermissions{
			Role:        role,
			Permissions: GetRolePermissions(role),
		})
	}

	return rolePerms
}
