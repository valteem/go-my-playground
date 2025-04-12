package storage

import (
	"context"
)

type Feature interface {
	CreateFeature() error
}

type FeatureSet interface {
	// CreateProductFeature creates new FeatureSet, adding provided Features (if any)
	CreateFeatureSet(ctx context.Context, features ...Feature) error
	// AddProductFeature adds provided Feature(s) to FeatureSet
	AddProductFeature(ctx context.Context, features ...Feature) error
}

type Product interface {
	// CreateProduct creates a new Product based on provided FeatureSet
	// Returns ID of new Product, or 0 if operation fails
	CreateProduct(ctx context.Context, set FeatureSet) (int, error)
	// UpdateProduct updates existing Product with provided FeatureSet
	// Returns error if Product with given ID does not exist
	UpdateProduct(ctx context.Context, id int, set FeatureSet) error
}
