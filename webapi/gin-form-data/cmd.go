package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.GET("/person", GetPerson)
	//	g.GET("/address", GetAddress)
	g.Run()
}
