package services

import (
	"context"

	"webapi/product-catalog/repository"
)

type ProductService struct {
	ProductRepository repository.Product
}

func (s *ProductService) CreateProduct(ctx context.Context, description string, fs repository.FeatureSet) (int, error) {
	id, err := s.ProductRepository.CreateProduct(ctx, description, fs)
	return id, err
}

func (s *ProductService) UpdateProduct(ctx context.Context, id int, set repository.FeatureSet) error {
	return s.ProductRepository.UpdateProduct(ctx, id, set)
}
