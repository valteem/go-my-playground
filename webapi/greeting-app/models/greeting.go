package models

import "time"

type Greeting struct {
	ID        int64     `json:"id"`
	Message   string    `json:"message"`
	CreatedBy int64     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

type AddGreetingRequest struct {
	Message string `json:"message" binding:"required"`
}
