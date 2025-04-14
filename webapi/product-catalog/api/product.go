package api

import (
	"github.com/gin-gonic/gin"

	"webapi/product-catalog/services"
)

// Used to inject service info into handler functions
type productRoutes struct {
	Services services.Services
}

func newProductRoutes(g *gin.RouterGroup, productService services.Services) {

	pr := &productRoutes{
		Services: productService,
	}

	g.POST("/create", pr.create)
}

func (pr *productRoutes) create(c *gin.Context) {
	// stub
	// pr.Services.Product.CreateProduct()
}
