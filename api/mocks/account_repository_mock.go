package mocks

import (
	"github.com/adisnuhic/go-gin-ntier/ecode"
	"github.com/adisnuhic/go-gin-ntier/models"
	apperror "github.com/adisnuhic/go-gin-ntier/pkg"
	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	mock mock.Mock
}

func NewAccountRepositoryMock() *AccountRepositoryMock {
	return &AccountRepositoryMock{}
}

func (m *AccountRepositoryMock) Register(user *models.User) (*models.User, *apperror.AppError) {
	args := m.mock.Called(user)

	usr := args.Get(0)
	err := args.Error(1)

	if err != nil {
		return nil, apperror.New(ecode.ErrUnableToCreateUserCode, err.Error(), ecode.ErrUnableToCreateUserMsg)
	}

	return usr.(*models.User), nil
}
