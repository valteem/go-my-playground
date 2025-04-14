package impl

import (
	"context"
	"fmt"

	"webapi/product-catalog/sqldb"
)

// TODO: add back implementation check (commented out for now to avoid import cycle)
// var _ repository.Product = (*ProductRepository)(nil)

// ProductRepository implements repository.Product interface
type ProductRepository struct {
	*sqldb.PostgresDB
}

func NewProductRepository(pg *sqldb.PostgresDB) *ProductRepository {
	return &ProductRepository{pg}
}

func (ps *ProductRepository) CreateProduct(ctx context.Context, description string) (int, error) {

	rows := ps.Pool.QueryRow(ctx, "insert into product (description) values ($1) returning id", description)

	var id int
	if err := rows.Scan(id); err != nil {
		return 0, fmt.Errorf("failed to add %s: %w", description, err)
	}

	return id, nil

}

func (ps *ProductRepository) UpdateProduct(ctx context.Context, id int) error {
	// add stub to comply with Product interface
	return nil
}
