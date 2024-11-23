package authservice

import (
	"ecommers/model"
	"ecommers/repository"
	"errors"

	"go.uber.org/zap"
)

type AuthService struct {
	Repo   repository.AllRepository
	Logger *zap.Logger
}

func NewAuthService(repo repository.AllRepository, Log *zap.Logger) AuthService {
	return AuthService{
		Repo:   repo,
		Logger: Log,
	}
}

func (s *AuthService) Login(login *model.Login) error {

	err := s.Repo.AuthRepo.Login(login)
	if err != nil {
		return errors.New("login failed: " + err.Error())
	}

	return nil
}
