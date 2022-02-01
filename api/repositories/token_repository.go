package repositories

import (
	"github.com/adisnuhic/hearken/db"
	"github.com/adisnuhic/hearken/ecode"
	"github.com/adisnuhic/hearken/models"
	apperror "github.com/adisnuhic/hearken/pkg"
)

// ITokenRepository interface
type ITokenRepository interface {
	CreateToken(token *models.Token) (*models.Token, *apperror.AppError)
	GetByToken(token string) (*models.Token, *apperror.AppError)
}

type tokenRepository struct {
	Store db.Store
}

// NewTokenRepository -
func NewTokenRepository(store db.Store) ITokenRepository {
	return &tokenRepository{
		Store: store,
	}
}

// CreateToken creates token
func (repo tokenRepository) CreateToken(token *models.Token) (*models.Token, *apperror.AppError) {
	if err := repo.Store.Create(token).Error; err != nil {
		return nil, apperror.New(ecode.ErrUnableToCreateTokenCode, err.Error(), ecode.ErrUnableToCreateTokenMsg)
	}

	return token, nil
}

// GetByToken returns token for provided token string
func (repo tokenRepository) GetByToken(token string) (*models.Token, *apperror.AppError) {
	model := new(models.Token)

	tx := repo.Store.Where("token = ?", token).Find(&model)

	if tx.Error != nil {
		return nil, apperror.New(ecode.ErrUnableToGetTokenCode, tx.Error.Error(), ecode.ErrUnableToGetTokenMsg)
	}

	return model, nil
}
