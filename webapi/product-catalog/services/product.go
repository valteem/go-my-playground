package services

import (
	"context"

	"webapi/product-catalog/model"
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

func (s *ProductService) CreateProduct(ctx context.Context, p *model.Product) (int, error) {
	id, err := s.ProductRepository.CreateProduct(ctx, p)
	return id, err
}

func (s *ProductService) UpdateProduct(ctx context.Context, p *model.Product) error {
	return s.ProductRepository.UpdateProduct(ctx, p)
}

func (s *ProductService) GetProductById(ctx context.Context, id int) (*model.Product, error) {
	p, err := s.ProductRepository.GetProductById(ctx, id)
	return p, err
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int) (int, error) {
	id, err := s.ProductRepository.DeleteProduct(ctx, id)
	return id, err
}
