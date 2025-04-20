package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"webapi/product-catalog/model"
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

type responseProduct struct {
	Id int `json:"id"`
}

func (pr *productRoutes) create(c *gin.Context) {

	var input model.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := pr.Services.Product.CreateProduct(c.Request.Context(), &input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, responseProduct{Id: id})

}
