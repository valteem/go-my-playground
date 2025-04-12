package services

import (
	"context"

	"webapi/product-catalog/storage"
)

type ProductService struct {
	ProductStorage storage.Product
}

func (s *ProductService) CreateProduct(ctx context.Context, set storage.FeatureSet) (int, error) {
	id, err := s.ProductStorage.CreateProduct(ctx, set)
	return id, err
}

func (s *ProductService) UpdateProduct(ctx context.Context, id int, set storage.FeatureSet) error {
	return s.ProductStorage.UpdateProduct(ctx, id, set)
}
