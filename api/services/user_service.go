package services

import (
	"github.com/adisnuhic/hearken/models"
	apperror "github.com/adisnuhic/hearken/pkg"
	"github.com/adisnuhic/hearken/repositories"
)

// IUserService interface
type IUserService interface {
	GetByID(id uint64) (*models.User, *apperror.AppError)
	GetByEmail(email string) (*models.User, *apperror.AppError)
	Exists(email string) bool
}

type userService struct {
	Repository repositories.IUserRepository
}

// NewUserService -
func NewUserService(repo repositories.IUserRepository) IUserService {
	return &userService{
		Repository: repo,
	}
}

// GetByID returns user for provided ID
func (svc userService) GetByID(id uint64) (*models.User, *apperror.AppError) {
	return svc.Repository.GetByID(id)
}

// GetByEmail returns user for provided email
func (svc userService) GetByEmail(email string) (*models.User, *apperror.AppError) {
	return svc.Repository.GetByEmail(email)
}

// Exists returns boolean if user exists
func (svc userService) Exists(email string) bool {
	return svc.Repository.Exists(email)
}
