package main

import (
	"greeting-app/config"
	"greeting-app/database"
	"greeting-app/handlers"
	"greeting-app/middleware"
	"greeting-app/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	database.InitDB(cfg)

	// Set up Gin router
	router := gin.Default()

	// Public routes
	router.POST("/signup", func(c *gin.Context) {
		handlers.Signup(c, cfg)
	})

	router.POST("/signin", func(c *gin.Context) {
		handlers.Login(c, cfg)
	})

	router.GET("/permissions", handlers.GetRolePermissions)

	// Protected routes - all authenticated users can get greetings
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware(cfg))
	{
		authorized.GET("/greet", handlers.GetRandomGreeting)

		// Routes requiring specific permissions
		authorized.POST("/add",
			middleware.PermissionMiddleware(models.PermissionAddGreeting),
			handlers.AddGreeting)

		authorized.GET("/greetings",
			middleware.PermissionMiddleware(models.PermissionManageGreetings),
			handlers.GetAllGreetings)

		authorized.POST("/user/role",
			middleware.PermissionMiddleware(models.PermissionManageUsers),
			handlers.UpdateUserRole)
	}

	// Get port from environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(router.Run(":" + port))
}
