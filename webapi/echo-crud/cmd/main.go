package main

import (
	"github.com/labstack/echo/v4"

	"webapi/echo-crud/server"
)

func main() {
	e := echo.New()

	e.POST("/items", server.CreateItem)
	e.GET("/items/:id", server.GetItem)
	e.PUT("/items/:id", server.UpdateItem)
	e.DELETE("/items/:id", server.DeleteItem)

	e.Logger.Fatal(e.Start(":44567"))
}
