package api

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func newErrorResponse(c *gin.Context, errStatus int, msg string) {
	err := errors.New(msg)
	c.JSON(errStatus, err.Error())
}
