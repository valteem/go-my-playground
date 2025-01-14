// https://gin-gonic.com/docs/examples/graceful-restart-or-stop/

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	timeoutSec = 5
)

func main() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Gin server running ...")
	})

	server := http.Server{
		Addr:    ":3001",
		Handler: router.Handler(),
	}

	// Start server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to listen and serve: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	// waiting for quit signal
	<-quit
	log.Println(" Shutting server down ...")

	ctx, cancel := context.WithTimeout(context.Background(), timeoutSec*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shut dwon server: %v\n", err)
	}
	now := time.Now()

	<-ctx.Done()
	log.Printf("server shut down after timeout %f seconds", time.Since(now).Seconds())

}
