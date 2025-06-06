package impl

import (
	"context"
	"fmt"

	"webapi/product-catalog/model"
	"webapi/product-catalog/repository/repoerr"
	"webapi/product-catalog/sqldb"

	"github.com/jackc/pgx/v5"
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

func (pr *ProductRepository) CreateProduct(ctx context.Context, p *model.Product) (int, error) {

	rows := pr.Pool.QueryRow(ctx, "insert into product (description) values ($1) returning id", p.Description)

	var id int
	if err := rows.Scan(id); err != nil {
		return 0, fmt.Errorf("failed to add %s: %w", p.Description, err)
	}

	return id, nil

}

func (pr *ProductRepository) UpdateProduct(ctx context.Context, p *model.Product) error {
	// add stub to comply with Product interface
	return nil
}

func (pr *ProductRepository) GetProductById(ctx context.Context, id int) (*model.Product, error) {
	// stub
	p := &model.Product{}
	return p, nil
}

func (pr *ProductRepository) DeleteProduct(ctx context.Context, id int) (int, error) {

	var idDeleted int
	err := pr.Pool.QueryRow(ctx, "delete from product where id=($1) returning id", id).Scan(idDeleted)

	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, repoerr.ErrNotFound
		}
		return 0, err
	}

	return idDeleted, nil

}
