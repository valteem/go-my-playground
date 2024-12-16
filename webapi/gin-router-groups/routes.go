package main

import "github.com/gin-gonic/gin"

func getRoutes(router *gin.Engine) {

	v1 := router.Group("/v1")
	addItemRoutes(v1)
	addStoreRoutes(v1)

	v2 := router.Group("/v2")
	addItemRoutes(v2)
	addStoreRoutes(v2)
}
