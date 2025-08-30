package handlers

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"greeting-app/database"
	"greeting-app/models"

	"github.com/gin-gonic/gin"
)

func GetRandomGreeting(c *gin.Context) {
	// Get all greetings from database
	rows, err := database.DB.Query(context.Background(), "SELECT id, message FROM greetings")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch greetings"})
		return
	}
	defer rows.Close()

	var greetings []models.Greeting
	for rows.Next() {
		var greeting models.Greeting
		err := rows.Scan(&greeting.ID, &greeting.Message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan greeting"})
			return
		}
		greetings = append(greetings, greeting)
	}

	if len(greetings) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No greetings found"})
		return
	}

	// Seed random generator
	rand.Seed(time.Now().UnixNano())
	randomGreeting := greetings[rand.Intn(len(greetings))]

	c.JSON(http.StatusOK, gin.H{
		"message": randomGreeting.Message,
	})
}

func AddGreeting(c *gin.Context) {
	userID := c.GetInt64("user_id")

	var req models.AddGreetingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert greeting into database
	var greeting models.Greeting
	err := database.DB.QueryRow(
		context.Background(),
		"INSERT INTO greetings (message, created_by) VALUES ($1, $2) RETURNING id, message, created_by, created_at",
		req.Message, userID,
	).Scan(&greeting.ID, &greeting.Message, &greeting.CreatedBy, &greeting.CreatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add greeting"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Greeting added successfully",
		"greeting": greeting,
	})
}

func GetAllGreetings(c *gin.Context) {
	rows, err := database.DB.Query(context.Background(), "SELECT id, message, created_by, created_at FROM greetings ORDER BY created_at DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch greetings"})
		return
	}
	defer rows.Close()

	var greetings []models.Greeting
	for rows.Next() {
		var greeting models.Greeting
		err := rows.Scan(&greeting.ID, &greeting.Message, &greeting.CreatedBy, &greeting.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan greeting"})
			return
		}
		greetings = append(greetings, greeting)
	}

	c.JSON(http.StatusOK, gin.H{
		"greetings": greetings,
		"count":     len(greetings),
	})
}
