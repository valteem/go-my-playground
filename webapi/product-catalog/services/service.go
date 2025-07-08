package services

import (
	"context"
	"time"

	"webapi/product-catalog/hashing"
	"webapi/product-catalog/model"
	"webapi/product-catalog/repository"
)

// Repeating repository.Product interface here enables injecting services to api layer
// without having to inject repository level objects too
// type Product interface {
// 	CreateProduct(ctx context.Context, description string) (int, error)
// 	UpdateProduct(ctx context.Context, id int) error
// 	GetProductById(ctx context.Context, id int) (*model.Product, error)
// }

type Product interface {
	CreateProduct(ctx context.Context, p *model.Product) (int, error)
	UpdateProduct(ctx context.Context, p *model.Product) error
	GetProductById(ctx context.Context, id int) (*model.Product, error)
	DeleteProduct(ctx context.Context, id int) (int, error)
}

type User interface {
	CreateUser(ctx context.Context, input model.User) (int, error)
	GetUserById(ctx context.Context, id int) (*model.User, error)
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	GetUserByNameAndPassword(ctx context.Context, name, password string) (*model.User, error)
}

// Product, User interfaces enable decoupling API and Repository layers
type Services struct {
	Product
	User
}

type ServiceDependencies struct {
	Repositories *repository.Repositories
	Hasher       hashing.Hasher

	SignKey  string
	TokenTTL time.Duration
}

func NewServices(deps *ServiceDependencies) *Services {
	return &Services{
		Product: NewProductService(deps.Repositories.Product),
		User:    NewUserService(deps.Repositories.User, deps.Hasher, deps.SignKey, deps.TokenTTL),
	}
}
