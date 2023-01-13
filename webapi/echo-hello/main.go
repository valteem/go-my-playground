package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

func main() {

	e := echo.New();
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":3000"))
	
}