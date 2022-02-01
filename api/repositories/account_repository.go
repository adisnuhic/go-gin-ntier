package repositories

import (
	"github.com/adisnuhic/hearken/db"
	"github.com/adisnuhic/hearken/ecode"
	"github.com/adisnuhic/hearken/models"
	apperror "github.com/adisnuhic/hearken/pkg"
)

// IAccountRepository interface
type IAccountRepository interface {
	Register(user *models.User) (*models.User, *apperror.AppError)
}

type accountRepository struct {
	Store db.Store
}

// NewAccountRepository -
func NewAccountRepository(store db.Store) IAccountRepository {
	return &accountRepository{
		Store: store,
	}
}

// Register user
func (repo accountRepository) Register(user *models.User) (*models.User, *apperror.AppError) {

	if err := repo.Store.Create(&user).Error; err != nil {
		return nil, apperror.New(ecode.ErrUnableToCreateUserCode, err.Error(), ecode.ErrUnableToCreateUserMsg)
	}

	return user, nil
}
