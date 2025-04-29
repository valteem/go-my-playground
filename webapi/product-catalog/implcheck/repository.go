package implcheck

import (
	"webapi/product-catalog/repository"
	"webapi/product-catalog/repository/impl"
)

var _ repository.Product = (*impl.ProductRepository)(nil)
var _ repository.User = (*impl.UserRepository)(nil)
