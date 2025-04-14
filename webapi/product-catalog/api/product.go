package api

import (
	"github.com/gin-gonic/gin"

	"webapi/product-catalog/services"
)

// Used to inject service info into handler functions
type productRoutes struct {
	productService services.Product
}

func newProductRoutes(g *gin.RouterGroup, productService services.Product) {

	pr := &productRoutes{
		productService: productService,
	}

	g.POST("/create", pr.create)
}

func (pr *productRoutes) create(c *gin.Context) {
	//stub
}
