package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"webapi/product-catalog/model"
	"webapi/product-catalog/services"
)

// Used to inject service info into handler functions
type productRoutes struct {
	Services *services.Services
}

func newProductRoutes(g *gin.RouterGroup, srv *services.Services) {

	pr := &productRoutes{
		Services: srv,
	}

	g.POST("/create", pr.create)
	g.DELETE("/delete", pr.delete)
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

	type response struct {
		Id int `json:"id"`
	}

	c.JSON(http.StatusCreated, response{Id: id})

}

type productDelete struct {
	Id int `json:"id"`
}

// https://stackoverflow.com/questions/4088350/is-rest-delete-really-idempotent
func (pr *productRoutes) delete(c *gin.Context) {

	var product productDelete

	if err := c.ShouldBindJSON(&product); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	_, err := pr.Services.Product.DeleteProduct(c.Request.Context(), product.Id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "product not found")
	}

}
