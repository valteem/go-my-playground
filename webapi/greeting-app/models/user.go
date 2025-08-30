package models

import "time"

// Roles
const (
	RoleAdmin            = "admin"
	RoleGreetingsManager = "greetings_manager"
	RoleUser             = "user"
)

// Permissions
const (
	PermissionAddGreeting     = "greeting:add"
	PermissionManageUsers     = "user:manage"
	PermissionGetGreeting     = "greeting:get"
	PermissionManageGreetings = "greeting:manage"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"` // "admin", "greetings_manager", or "user"
	CreatedAt time.Time `json:"created_at"`
}

type UserResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserRoleUpdateRequest struct {
	UserID int64  `json:"user_id" binding:"required"`
	Role   string `json:"role" binding:"required,oneof=admin greetings_manager user"`
}

type Permission struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RolePermissions struct {
	Role        string       `json:"role"`
	Permissions []Permission `json:"permissions"`
}
