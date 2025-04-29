package implcheck

import (
	"webapi/product-catalog/services"
)

var _ services.Product = (*services.ProductService)(nil)
var _ services.User = (*services.UserService)(nil)
