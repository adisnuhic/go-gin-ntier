package models

import "time"

// User -
type User struct {
	ID          uint64    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email" binding:"required"`
	IsConfirmed bool      `json:"is_confirmed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Users collection
type Users []User
