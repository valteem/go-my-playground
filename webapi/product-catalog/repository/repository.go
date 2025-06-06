package repository

import (
	"context"

	"webapi/product-catalog/model"
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
	// CreateProduct creates a new Product based on provided input.
	// Ignores model.Product.Id if set.
	// Returns Id of new Product, or 0 if operation fails
	CreateProduct(ctx context.Context, p *model.Product) (int, error)
	// UpdateProduct updates existing Product with provided input
	// Returns error if Product with given *model.Product.Id does not exist
	UpdateProduct(ctx context.Context, p *model.Product) error
	// GetProductById returns Product by provided id, returns error if Product is no found
	GetProductById(ctx context.Context, id int) (*model.Product, error)
	// DeleteProduct attempts to delete Product record
	// Returns id of deleted product and nil if the Product with given the id exists,
	// 0 and errNotFound if no Product exist with the given id
	DeleteProduct(ctx context.Context, id int) (int, error)
}

type User interface {
	CreateUser(ctx context.Context, input model.User) (int, error)
	GetUserById(ctx context.Context, id int) (*model.User, error)
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	GetUserByNameAndPassword(ctx context.Context, name, password string) (*model.User, error)
}

type Repositories struct {
	Product
	User
}

// Link to database layer
func NewRepositories(pg *sqldb.PostgresDB) *Repositories {
	return &Repositories{
		Product: impl.NewProductRepository(pg),
		User:    impl.NewUserRepository(pg),
	}
}
