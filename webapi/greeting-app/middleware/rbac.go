package middleware

import (
	"net/http"

	"greeting-app/database"
	//	"greeting-app/models"
	"greeting-app/rbac"

	"github.com/gin-gonic/gin"
)

// PermissionMiddleware checks if the user has the required permission
func PermissionMiddleware(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("user_id")

		// Get user role
		var userRole string
		err := database.DB.QueryRow(
			c,
			"SELECT role FROM users WHERE id = $1",
			userID,
		).Scan(&userRole)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify user permissions"})
			c.Abort()
			return
		}

		// Check if user has required permission
		if !rbac.HasPermission(userRole, permission) {
			c.JSON(http.StatusForbidden, gin.H{
				"error":               "Insufficient permissions",
				"required_permission": permission,
				"user_role":           userRole,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// RoleMiddleware checks if the user has one of the specified roles
func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("user_id")

		// Get user role
		var userRole string
		err := database.DB.QueryRow(
			c,
			"SELECT role FROM users WHERE id = $1",
			userID,
		).Scan(&userRole)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify user role"})
			c.Abort()
			return
		}

		// Check if user has one of the allowed roles
		allowed := false
		for _, role := range allowedRoles {
			if userRole == role {
				allowed = true
				break
			}
		}

		if !allowed {
			c.JSON(http.StatusForbidden, gin.H{
				"error":          "Insufficient role privileges",
				"required_roles": allowedRoles,
				"user_role":      userRole,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
