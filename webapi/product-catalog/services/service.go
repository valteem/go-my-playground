package services

import (
	"context"

	"webapi/product-catalog/repository"
)

// Repeating repository.Product interface here enables injecting services to api layer
// without having to inject repository level objects too
type Product interface {
	CreateProduct(ctx context.Context, description string) (int, error)
	UpdateProduct(ctx context.Context, id int) error
}

type Services struct {
	Product
}

type ServiceDependencies struct {
	Repositories *repository.Repositories
}

func NewServices(deps *ServiceDependencies) *Services {
	return &Services{
		Product: NewProductService(deps.Repositories.Product),
	}
}
