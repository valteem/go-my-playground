package handlers

import (
	"context"
	"greeting-app/config"
	"greeting-app/database"
	"greeting-app/models"
	"greeting-app/rbac"
	"greeting-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context, cfg *config.Config) {
	var req models.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Insert user into database with default 'user' role
	var user models.User
	err = database.DB.QueryRow(
		context.Background(),
		"INSERT INTO users (username, password, role) VALUES ($1, $2, $3) RETURNING id, username, role, created_at",
		req.Username, string(hashedPassword), models.RoleUser,
	).Scan(&user.ID, &user.Username, &user.Role, &user.CreatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Username already exists or database error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": models.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		},
	})
}

func Login(c *gin.Context, cfg *config.Config) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user from database
	var user models.User
	err := database.DB.QueryRow(
		context.Background(),
		"SELECT id, username, password, role FROM users WHERE username = $1",
		req.Username,
	).Scan(&user.ID, &user.Username, &user.Password, &user.Role)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Username, cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user": models.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		},
	})
}

func UpdateUserRole(c *gin.Context) {
	// Check if requesting user has user management permission
	requestingUserID := c.GetInt64("user_id")

	var requestingUserRole string
	err := database.DB.QueryRow(
		c,
		"SELECT role FROM users WHERE id = $1",
		requestingUserID,
	).Scan(&requestingUserRole)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify user permissions"})
		return
	}

	if !rbac.HasPermission(requestingUserRole, models.PermissionManageUsers) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions to manage users"})
		return
	}

	// Parse request
	var req models.UserRoleUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate role
	validRoles := []string{models.RoleAdmin, models.RoleGreetingsManager, models.RoleUser}
	valid := false
	for _, role := range validRoles {
		if req.Role == role {
			valid = true
			break
		}
	}

	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role specified"})
		return
	}

	// Prevent users from demoting themselves from admin
	if requestingUserID == req.UserID && requestingUserRole == models.RoleAdmin && req.Role != models.RoleAdmin {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot demote yourself from admin"})
		return
	}

	// Update user role
	_, err = database.DB.Exec(
		context.Background(),
		"UPDATE users SET role = $1 WHERE id = $2",
		req.Role, req.UserID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User role updated successfully",
		"user_id": req.UserID,
		"role":    req.Role,
	})
}

func GetRolePermissions(c *gin.Context) {
	rolePermissions := rbac.GetAllRolePermissions()
	c.JSON(http.StatusOK, gin.H{
		"role_permissions": rolePermissions,
	})
}
