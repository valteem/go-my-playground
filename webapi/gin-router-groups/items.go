package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addItemRoutes(rg *gin.RouterGroup) {

	items := rg.Group("/items")

	items.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "items")
	})

	items.GET("/descriptions", func(c *gin.Context) {
		c.String(http.StatusOK, `descriptions`)
	})

	items.GET("/images", func(c *gin.Context) {
		c.String(http.StatusOK, `images`)
	})

}
