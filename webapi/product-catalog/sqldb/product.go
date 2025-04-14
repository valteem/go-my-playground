package sqldb

import (
	"context"
	"fmt"

	"webapi/product-catalog/repository"
)

var _ repository.Product = (*ProductRepository)(nil)

// ProductStorage implements storage.Product interface
type ProductRepository struct {
	*PostgresDB
}

func NewProductRepository(pg *PostgresDB) *ProductRepository {
	return &ProductRepository{pg}
}

func (ps *ProductRepository) CreateProduct(ctx context.Context, description string, fs repository.FeatureSet) (int, error) {

	rows := ps.Pool.QueryRow(ctx, "insert into product (description) values ($1) returning id", description)

	var id int
	if err := rows.Scan(id); err != nil {
		return 0, fmt.Errorf("failed to add %s: %w", description, err)
	}

	return id, nil

}

func (ps *ProductRepository) UpdateProduct(ctx context.Context, id int, fs repository.FeatureSet) error {
	// add stub to comply with Product interface
	return nil
}
