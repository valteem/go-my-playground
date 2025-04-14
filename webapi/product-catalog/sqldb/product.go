package sqldb

import (
	"context"
	"fmt"

	"webapi/product-catalog/repository"
)

var _ repository.Product = (*ProductStorage)(nil)

// ProductStorage implements storage.Product interface
type ProductStorage struct {
	*PostgresDB
}

func NewProductStorage(pg *PostgresDB) *ProductStorage {
	return &ProductStorage{pg}
}

func (ps *ProductStorage) CreateProduct(ctx context.Context, description string, fs repository.FeatureSet) (int, error) {

	rows := ps.Pool.QueryRow(ctx, "insert into product (description) values ($1) returning id", description)

	var id int
	if err := rows.Scan(id); err != nil {
		return 0, fmt.Errorf("failed to add %s: %w", description, err)
	}

	return id, nil

}

func (ps *ProductStorage) UpdateProduct(ctx context.Context, id int, fs repository.FeatureSet) error {
	// add stub to comply with Product interface
	return nil
}
