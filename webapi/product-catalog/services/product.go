package services

import (
	"context"

	"webapi/product-catalog/storage"
)

type ProductService struct {
	ProductStorage storage.Product
}

func (s *ProductService) CreateProduct(ctx context.Context, description string, fs storage.FeatureSet) (int, error) {
	id, err := s.ProductStorage.CreateProduct(ctx, description, fs)
	return id, err
}

func (s *ProductService) UpdateProduct(ctx context.Context, id int, set storage.FeatureSet) error {
	return s.ProductStorage.UpdateProduct(ctx, id, set)
}
