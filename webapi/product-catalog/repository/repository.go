package repository

import (
	"context"

	"webapi/product-catalog/repository/impl"
	"webapi/product-catalog/sqldb"
)

// TODO: add back Feature/FeatureSet

/* type Feature interface {
	CreateFeature() error
}

type FeatureSet interface {
	// CreateProductFeature creates new FeatureSet, adding provided Features (if any)
	CreateFeatureSet(ctx context.Context, features ...Feature) error
	// AddProductFeature adds provided Feature(s) to FeatureSet
	AddProductFeature(ctx context.Context, features ...Feature) error
} */

type Product interface {
	// CreateProduct creates a new Product based on provided FeatureSet
	// Returns ID of new Product, or 0 if operation fails
	CreateProduct(ctx context.Context, description string) (int, error)
	// UpdateProduct updates existing Product with provided FeatureSet
	// Returns error if Product with given ID does not exist
	UpdateProduct(ctx context.Context, id int) error
}

type Repositories struct {
	Product
}

// Link to database layer
func NewRepositories(pg *sqldb.PostgresDB) *Repositories {
	return &Repositories{
		Product: impl.NewProductRepository(pg),
	}
}
