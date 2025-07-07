package tree

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func responseWithMsg(c *gin.Context, msg string) {
	c.String(http.StatusOK, msg)
}

func buildServer() *gin.Engine {

	engine := gin.New()

	engine.GET("/", func(c *gin.Context) {
		responseWithMsg(c, "/")
	})

	engine.GET("/account", func(c *gin.Context) {
		responseWithMsg(c, "/account")
	})
	engine.GET("/account/active", func(c *gin.Context) {
		responseWithMsg(c, "/account/active")
	})
	engine.GET("/account/retired", func(c *gin.Context) {
		responseWithMsg(c, "/account/retired")
	})

	engine.GET("/product", func(c *gin.Context) {
		responseWithMsg(c, "/product")
	})
	engine.GET("/product/active", func(c *gin.Context) {
		responseWithMsg(c, "/product/active")
	})
	engine.GET("/product/retired", func(c *gin.Context) {
		responseWithMsg(c, "/product/retired")
	})

	engine.POST("/account", func(c *gin.Context) {
		responseWithMsg(c, "new account")
	})

	engine.POST("/product", func(c *gin.Context) {
		responseWithMsg(c, "new product")
	})

	return engine

}

func runServer(port string) {

	server := buildServer()

	if err := server.Run(port); err != nil {
		log.Fatalf("failed to serve on %s: %v", port, err)
	}

}
