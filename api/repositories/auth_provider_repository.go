package repositories

import (
	"github.com/adisnuhic/go-gin-ntier/db"
	"github.com/adisnuhic/go-gin-ntier/ecode"
	"github.com/adisnuhic/go-gin-ntier/models"
	apperror "github.com/adisnuhic/go-gin-ntier/pkg"
)

// IAuthProviderRepository interface
type IAuthProviderRepository interface {
	GetByUserID(id uint64) (*models.AuthProvider, *apperror.AppError)
	Update(auth *models.AuthProvider) *apperror.AppError
	Create(auth *models.AuthProvider) *apperror.AppError
	Delete(auth *models.AuthProvider) *apperror.AppError
	Save(auth *models.AuthProvider) *apperror.AppError
}

type authProviderRepository struct {
	Store db.Store
}

// NewAuthProviderRepository -
func NewAuthProviderRepository(store db.Store) IAuthProviderRepository {
	return &authProviderRepository{
		Store: store,
	}
}

// GetByUserID return auth provider for provided ID
func (repo authProviderRepository) GetByUserID(id uint64) (*models.AuthProvider, *apperror.AppError) {
	model := &models.AuthProvider{}

	if err := repo.Store.Where("user_id = ?", id).Find(model).Error; err != nil {
		return nil, apperror.New(ecode.ErrUnableToFetchAuthCode, err.Error(), ecode.ErrUnableToFetchAuthMsg)
	}
	return model, nil
}

// Update saves auth data
func (repo authProviderRepository) Save(auth *models.AuthProvider) *apperror.AppError {
	if auth.UserID > 0 {
		return repo.Update(auth)
	}

	return repo.Create(auth)
}

// Update updates auth data
func (repo authProviderRepository) Update(auth *models.AuthProvider) *apperror.AppError {
	if err := repo.Store.Save(auth).Error; err != nil {
		return apperror.New(ecode.ErrUnableToSaveAuthCode, err.Error(), ecode.ErrUnableToSaveAuthMsg)
	}
	return nil
}

// Create creates auth data
func (repo authProviderRepository) Create(auth *models.AuthProvider) *apperror.AppError {
	if err := repo.Store.Create(&auth).Error; err != nil {
		return apperror.New(ecode.ErrUnableToCreateAuthCode, err.Error(), ecode.ErrUnableToCreateAuthMsg)
	}

	return nil
}

// Deletes auth data
func (repo authProviderRepository) Delete(auth *models.AuthProvider) *apperror.AppError {
	return nil
}
