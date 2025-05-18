package main

import (
	"github.com/gin-gonic/gin"
)

const (
	userIdCtxKey = "userId"
)

func UserId() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Set(userIdCtxKey, "some-random-user") // TODO: add extracting userId from JWT
	})
}
