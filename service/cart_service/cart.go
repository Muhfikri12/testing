package cartservice

import (
	"ecommers/model"
	"ecommers/repository"

	"go.uber.org/zap"
)

type CartsService struct {
	Repo   repository.AllRepository
	Logger *zap.Logger
}

func NewCartsService(repo repository.AllRepository, Log *zap.Logger) CartsService {
	return CartsService{
		Repo:   repo,
		Logger: Log,
	}
}

func (cs *CartsService) TotalProducts(token string) (*[]model.Cart, error) {

	carts, err := cs.Repo.CartRepo.TotalCarts(token)
	if err != nil {
		cs.Logger.Error("Error from service: " + err.Error())
		return nil, err
	}

	return carts, nil
}
