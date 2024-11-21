package usersservice

import (
	"ecommers/model"
	"ecommers/repository"
	"errors"

	"go.uber.org/zap"
)

type UsersService struct {
	Repo   repository.AllRepository
	Logger *zap.Logger
}

func NewUsersService(repo repository.AllRepository, Log *zap.Logger) UsersService {
	return UsersService{
		Repo:   repo,
		Logger: Log,
	}
}

func (s *UsersService) Login(login *model.Login) error {

	err := s.Repo.UsersRepo.Login(login)
	if err != nil {
		return errors.New("login failed: " + err.Error())
	}

	return nil
}
