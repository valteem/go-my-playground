package services

import (
	"webapi/product-catalog/repository"
)

type Services struct {
	Product repository.Product
}

type ServiceDependencies struct {
	Repositories *repository.Repositories
}

func NewServices(deps *ServiceDependencies) *Services {
	return &Services{
		Product: NewProductService(deps.Repositories.Product),
	}
}
