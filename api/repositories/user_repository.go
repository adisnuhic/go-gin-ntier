package repositories

import (
	"github.com/adisnuhic/go-gin-ntier/db"
	"github.com/adisnuhic/go-gin-ntier/ecode"
	"github.com/adisnuhic/go-gin-ntier/models"
	apperror "github.com/adisnuhic/go-gin-ntier/pkg"
)

// IUserRepository interface
type IUserRepository interface {
	GetByID(id uint64) (*models.User, *apperror.AppError)
	GetByEmail(email string) (*models.User, *apperror.AppError)
	Exists(email string) bool
}

type userRepository struct {
	Store db.Store
}

// NewUserRepository -
func NewUserRepository(store db.Store) IUserRepository {
	return &userRepository{
		Store: store,
	}
}

// GetByID returns user for provided ID
func (repo userRepository) GetByID(id uint64) (*models.User, *apperror.AppError) {
	model := new(models.User)
	if err := repo.Store.Where("id = ?", id).Find(&model).Error; err != nil {
		return nil, apperror.New(ecode.ErrUnableToFetchUserCode, err.Error(), ecode.ErrUnableToFetchUserMsg)
	}

	return model, nil
}

// GetByEmail returns user for provided email
func (repo userRepository) GetByEmail(email string) (*models.User, *apperror.AppError) {
	model := new(models.User)

	if err := repo.Store.Where("email = ?", email).Find(&model).Error; err != nil {
		return nil, apperror.New(ecode.ErrUnableToFetchUserCode, err.Error(), ecode.ErrUnableToFetchUserMsg)
	}
	return model, nil
}

// Exists returns boolean if user exists
func (repo userRepository) Exists(email string) bool {
	model := new(models.User)
	repo.Store.Where("email = ?", email).Find(&model)

	if model.ID > 0 {
		return true
	}

	return false
}
