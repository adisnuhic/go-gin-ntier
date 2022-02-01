package services

import (
	"github.com/adisnuhic/hearken/models"
	apperror "github.com/adisnuhic/hearken/pkg"
	"github.com/adisnuhic/hearken/repositories"
)

// IAccountService interface
type IAccountService interface {
	Register(user *models.User) (*models.User, *apperror.AppError)
}

type accountService struct {
	Repository repositories.IAccountRepository
}

// NewAccountService -
func NewAccountService(repo repositories.IAccountRepository) IAccountService {
	return &accountService{
		Repository: repo,
	}
}

// Register user
func (svc accountService) Register(user *models.User) (*models.User, *apperror.AppError) {
	return svc.Repository.Register(user)
}
