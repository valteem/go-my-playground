package api

import (
	"net/http"

	"webapi/product-catalog/model"
	"webapi/product-catalog/services"

	"github.com/gin-gonic/gin"
)

type userRoutes struct {
	services *services.Services
}

func newUserRoutes(g *gin.RouterGroup, srv *services.Services) {

	ur := &userRoutes{services: srv}

	g.POST("/signup", ur.signup)

}

func (ur *userRoutes) signup(c *gin.Context) {

	var input model.User

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return

	}

	id, err := ur.services.User.CreateUser(c.Request.Context(), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	type response struct {
		Id int `json:"id"`
	}

	c.JSON(http.StatusCreated, response{Id: id})

}
