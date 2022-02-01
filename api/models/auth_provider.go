package models

import "time"

// AuthProvider -
type AuthProvider struct {
	Provider  string    `json:"provider"`
	UserID    uint64    `json:"user_id" binding:"required"`
	UID       string    `json:"uid" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AuthProviders collection
type AuthProviders []AuthProvider
