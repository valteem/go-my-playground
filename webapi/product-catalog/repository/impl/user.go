package impl

import (
	"context"
	"errors"
	"fmt"

	"webapi/product-catalog/model"
	"webapi/product-catalog/repository/repoerr"
	"webapi/product-catalog/sqldb"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepository struct {
	*sqldb.PostgresDB
}

func NewUserRepository(pg *sqldb.PostgresDB) *UserRepository {
	return &UserRepository{pg}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user model.User) (int, error) {

	rows := ur.Pool.QueryRow(ctx, "insert into users (name, password) values ($1, $2) returning id", user.Name, user.Password)

	var id int
	if err := rows.Scan(id); err != nil {
		var pgErr *pgconn.PgError
		// errors.As: 2nd argument is set to the match found (if any)
		if ok := errors.As(err, &pgErr); ok {
			// Class 23 — Integrity Constraint Violation
			// 23505	unique_violation
			if pgErr.Code == "23505" {
				return 0, repoerr.ErrAlreadyExists
			}
		}
		return 0, fmt.Errorf("failed to add %s: %w", user.Name, err)
	}

	return id, nil

}

func (ur *UserRepository) GetUserByNameAndPassword(ctx context.Context, name, password string) (*model.User, error) {

	user := model.User{}

	rows, err := ur.Pool.Query(ctx, "select id, name, password, created_at from users where name = $1 and password = $2", name, password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &model.User{}, repoerr.ErrNotFound
		}
		return nil, err
	}

	err = rows.Scan(&user.Id, &user.Name, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (ur *UserRepository) GetUserById(ctx context.Context, id int) (*model.User, error) {

	user := model.User{}

	rows, err := ur.Pool.Query(ctx, "select id, name, password, created_at from users where id = $1", id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &model.User{}, repoerr.ErrNotFound
		}
		return nil, err
	}

	err = rows.Scan(&user.Id, &user.Name, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (ur *UserRepository) GetUserByName(ctx context.Context, name string) (*model.User, error) {

	user := model.User{}

	rows, err := ur.Pool.Query(ctx, "select id, name, password, created_at from users where name = $1", name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &model.User{}, repoerr.ErrNotFound
		}
		return nil, err
	}

	err = rows.Scan(&user.Id, &user.Name, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil

}
