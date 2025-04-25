package impl

import (
	"context"
	"fmt"

	"webapi/product-catalog/model"
	"webapi/product-catalog/sqldb"
)

type UserRepository struct {
	*sqldb.PostgresDB
}

func NewUserRepository(pg *sqldb.PostgresDB) *UserRepository {
	return &UserRepository{pg}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user model.User) (int, error) {

	rows := ur.Pool.QueryRow(ctx, "insert into user (name, password) values ($1, $2) returning id", user.Name, user.Password)

	var id int
	if err := rows.Scan(id); err != nil {
		return 0, fmt.Errorf("failed to add %s: %w", user.Name, err)
	}

	return id, nil

}
