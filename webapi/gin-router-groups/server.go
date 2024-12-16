package main

import (
	"github.com/gin-gonic/gin"
)

func runServer(e *gin.Engine, port string) {
	getRoutes(e)
	e.Run(port)
}
