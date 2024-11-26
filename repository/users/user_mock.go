package users

import (
	"ecommers/model"

	"github.com/stretchr/testify/mock"
)

type UsersRepositoryMock struct {
	mock.Mock
}

func (usersRepositoryMock *UsersRepositoryMock) GetDetailUser(token string) (*model.Users, error) {
	return nil, nil
}
