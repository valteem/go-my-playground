package services

import (
	"context"

	"webapi/product-catalog/repository"
)

type ProductService struct {
	ProductRepository repository.Product
}

func NewProductService(r repository.Product) *ProductService {
	return &ProductService{
		ProductRepository: r,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, description string) (int, error) {
	id, err := s.ProductRepository.CreateProduct(ctx, description)
	return id, err
}

func (s *ProductService) UpdateProduct(ctx context.Context, id int) error {
	return s.ProductRepository.UpdateProduct(ctx, id)
}
