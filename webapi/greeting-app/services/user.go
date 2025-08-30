package services

import (
	"context"

	"greeting-app/database"
	"greeting-app/models"
)

func GetUserByID(userID int64) (*models.User, error) {
	var user models.User
	err := database.DB.QueryRow(
		context.Background(),
		"SELECT id, username, password, role, created_at FROM users WHERE id = $1",
		userID,
	).Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUserRole(userID int64, role string) error {
	_, err := database.DB.Exec(
		context.Background(),
		"UPDATE users SET role = $1 WHERE id = $2",
		role, userID,
	)
	return err
}

func IsUserAdmin(userID int64) (bool, error) {
	var role string
	err := database.DB.QueryRow(
		context.Background(),
		"SELECT role FROM users WHERE id = $1",
		userID,
	).Scan(&role)

	if err != nil {
		return false, err
	}

	return role == "admin", nil
}
