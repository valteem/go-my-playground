package api

import (
	"github.com/gin-gonic/gin"

	"webapi/product-catalog/services"
)

func NewRouter(handler *gin.Engine, services services.Services) {

	api := handler.Group("/api")

	newProductRoutes(api, services)

}
