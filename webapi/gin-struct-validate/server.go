package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// go-playground/validator `validate` tags work as `binding` tags of Gin
type Person struct {
	Name  string `json:"name" binding:"alpha"`
	Email string `json:"email" binding:"email"`
}

func CreatePerson(c *gin.Context) {

	p := &Person{}

	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, p)

}

func runServer(port string) {

	e := gin.Default()

	e.PUT("/person", CreatePerson)

	if err := e.Run(port); err != nil {
		log.Fatalf("failed to start server: %v", err)
		return
	}

}
