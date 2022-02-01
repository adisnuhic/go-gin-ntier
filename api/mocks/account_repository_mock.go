package mocks

import (
	"github.com/adisnuhic/hearken/ecode"
	"github.com/adisnuhic/hearken/models"
	apperror "github.com/adisnuhic/hearken/pkg"
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
