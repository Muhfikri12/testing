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

func (cs *CartsService) GetDetailCart(token string) (*[]model.Products, error) {
	products, err := cs.Repo.CartRepo.GetDetailCart(token)
	if err != nil {
		cs.Logger.Error("Error from service: " + err.Error())
		return nil, err
	}

	return products, nil
}

func (cs *CartsService) AddItemToCart(token string, cart *model.Products) error {

	err := cs.Repo.CartRepo.AddItemToCart(token, cart)
	if err != nil {
		cs.Logger.Error("Error from service AddItemToCart: " + err.Error())
		return err
	}

	return nil
}

func (cs *CartsService) UpdateCart(token string, cart *model.Products) error {

	err := cs.Repo.CartRepo.UpdateCart(token, cart)
	if err != nil {
		cs.Logger.Error("Error from service UpdateCart: " + err.Error())
		return err
	}

	return nil
}

func (cs *CartsService) DeleteCart(token string, cart *model.Products) error {

	err := cs.Repo.CartRepo.DeleteCart(token, cart)
	if err != nil {
		cs.Logger.Error("Error from service DeleteCart: " + err.Error())
		return err
	}

	return nil
}