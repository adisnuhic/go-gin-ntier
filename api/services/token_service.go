package services

import (
	"time"

	"github.com/adisnuhic/go-gin-ntier/models"
	apperror "github.com/adisnuhic/go-gin-ntier/pkg"
	"github.com/adisnuhic/go-gin-ntier/repositories"
	"github.com/rs/xid"
)

const (
	// TokenTypeRefreshToken - refresh token
	TokenTypeRefreshToken = 1

	// RefreshTokenDuration holds duration value in minutes for refresh token
	RefreshTokenDuration = 43200 // 30 days
)

// ITokenService interface
type ITokenService interface {
	CreateRefreshToken(userID uint64, email string) (*models.Token, *apperror.AppError)
	GetByToken(token string) (*models.Token, *apperror.AppError)
}

type tokenService struct {
	Repository repositories.ITokenRepository
}

// NewTokenService -
func NewTokenService(repo repositories.ITokenRepository) ITokenService {
	return &tokenService{
		Repository: repo,
	}
}

// CreateRefreshToken creates refresh token in DB
func (svc tokenService) CreateRefreshToken(userID uint64, email string) (*models.Token, *apperror.AppError) {
	expireAt := time.Now().Add(time.Minute * time.Duration(RefreshTokenDuration))
	token := &models.Token{}
	token.UserID = userID
	token.TokenType = TokenTypeRefreshToken
	token.Token = xid.New().String()
	token.ExpiresAt = expireAt

	return svc.Repository.CreateToken(token)
}

func (svc tokenService) GetByToken(token string) (*models.Token, *apperror.AppError) {
	return svc.Repository.GetByToken(token)
}
