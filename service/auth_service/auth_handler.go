package authservice

import (
	"ecommers/model"
	"ecommers/repository"
	"errors"

	"go.uber.org/zap"
)

type AuthServiceInterface interface {
	Login(login *model.Login) error
	Register(user *model.Register) error
}

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

func (u *AuthService) Register(user *model.Register) error {

	if err := u.Repo.AuthRepo.Register(user); err != nil {
		u.Logger.Error("Error from Register Service: " + err.Error())
		return err
	}

	return nil
}
