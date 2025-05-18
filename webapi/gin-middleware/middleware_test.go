package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestUserId(t *testing.T) {

	r := gin.New()
	authGroup := r.Group("/auth")
	authGroup.Use(UserId())
	authGroup.POST("/login", func(c *gin.Context) {
		c.Header("User-Id", c.Keys[userIdCtxKey].(string))
	})

	go r.Run(":3001")
	time.Sleep(100 * time.Millisecond) // allow test server some time to sleep

	resp, err := http.Post("http://localhost:3001/auth/login", "application/json", nil)
	if err != nil {
		t.Fatalf("failed to get a response: %v", err)
	}

	if actual, expected := resp.Header["User-Id"][0], "some-random-user"; actual != expected {
		t.Errorf("User-Id: get %q, expect %q", actual, expected)
	}

}
