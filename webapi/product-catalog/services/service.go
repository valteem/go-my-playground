package services

import (
	//	"context"

	//	"webapi/product-catalog/model"
	"time"
	"webapi/product-catalog/hashing"
	"webapi/product-catalog/repository"
)

// Repeating repository.Product interface here enables injecting services to api layer
// without having to inject repository level objects too
// type Product interface {
// 	CreateProduct(ctx context.Context, description string) (int, error)
// 	UpdateProduct(ctx context.Context, id int) error
// 	GetProductById(ctx context.Context, id int) (*model.Product, error)
// }

type UserInput struct {
	Name     string
	Password string // plain password, before hashing
}

type Services struct {
	Product repository.Product
	User    repository.User
}

type ServiceDependencies struct {
	Repositories *repository.Repositories
	Hasher       *hashing.Hasher

	SignKey  string
	TokenTTL time.Duration
}

func NewServices(deps *ServiceDependencies) *Services {
	return &Services{
		Product: NewProductService(deps.Repositories.Product),
		User:    NewUserService(deps.Repositories.User, *deps.Hasher, deps.SignKey, deps.TokenTTL),
	}
}
