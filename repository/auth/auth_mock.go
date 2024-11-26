package auth

import (
	"ecommers/model"

	"github.com/stretchr/testify/mock"
)

type AuthRepositoryMock struct {
	mock.Mock
}

func (am *AuthRepositoryMock) Login(login *model.Login) error {

	return nil
}

func (am *AuthRepositoryMock) Register(user *model.Register) error {
	args := am.Called(user)
	if customerResult := args.Get(0); customerResult != nil {
		return args.Error(1)
	}
	return args.Error(1)
}
