package timeout

import (
	"log"
	"net/http"
	"time"

	gintimeout "github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

const (
	respDelay         = 100 * time.Millisecond
	quickHandlerDelay = 50 * time.Millisecond
	slowHandlerDelay  = 150 * time.Millisecond

	quickResponseMsg   = "quick response"
	slowResponseMsg    = "slow response"
	timeoutResponseMsg = "timeout response"
)

func responseHandler(c *gin.Context) {
	c.String(http.StatusRequestTimeout, timeoutResponseMsg)
}

func timeoutMiddleware() gin.HandlerFunc {
	return gintimeout.New(
		gintimeout.WithTimeout(respDelay),
		gintimeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		gintimeout.WithResponse(responseHandler),
	)
}

func runServer(port string) {

	engine := gin.New()

	engine.Use(timeoutMiddleware())

	engine.GET("/quick", func(c *gin.Context) {
		time.Sleep(quickHandlerDelay)
		c.String(http.StatusOK, quickResponseMsg)
	})

	engine.GET("/slow", func(c *gin.Context) {
		time.Sleep(slowHandlerDelay)
		c.String(http.StatusOK, timeoutResponseMsg)
	})

	if err := engine.Run(port); err != nil {
		log.Fatalf("failed to serve at %s: %v", port, err)
	}

}
