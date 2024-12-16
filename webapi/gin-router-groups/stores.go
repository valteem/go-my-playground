package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addStoreRoutes(rg *gin.RouterGroup) {

	stores := rg.Group("/stores")

	stores.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "stores")
	})

	stores.GET("/locations", func(c *gin.Context) {
		c.String(http.StatusOK, "locations")
	})

}
