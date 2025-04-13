package sqldb

import (
	"context"
	"fmt"

	"webapi/product-catalog/storage"
)

type ProductStorage struct {
	*PostgresDB
}

func NewProductStorage(pg *PostgresDB) *ProductStorage {
	return &ProductStorage{pg}
}

func (ps *ProductStorage) CreateProduct(ctx context.Context, description string, fs storage.FeatureSet) (int, error) {

	rows := ps.Pool.QueryRow(context.Background(), "insert into product (description) values ($1) returning id", description)

	var id int
	if err := rows.Scan(id); err != nil {
		return 0, fmt.Errorf("failed to add %s: %w", description, err)
	}

	return id, nil

}
