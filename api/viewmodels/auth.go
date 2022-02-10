package viewmodels

import "github.com/adisnuhic/go-gin-ntier/models"

// Auth -
type Auth struct {
	User         *models.User `json:"user"`
	Token        string       `json:"token"`
	RefreshToken string       `json:"refresh_token"`
}
