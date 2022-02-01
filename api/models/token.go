package models

import "time"

// Token -
type Token struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id" binding:"required"`
	Token     string    `json:"token" binding:"required"`
	TokenType uint64    `json:"token_type" binding:"required"`
	Meta      string    `json:"meta"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
