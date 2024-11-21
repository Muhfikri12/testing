package checkoutservice

import (
	"ecommers/model"
	"ecommers/repository"

	"go.uber.org/zap"
)

type CheckoutsService struct {
	Repo   repository.AllRepository
	Logger *zap.Logger
}

func NewCheckoutsService(repo repository.AllRepository, Log *zap.Logger) CheckoutsService {
	return CheckoutsService{
		Repo:   repo,
		Logger: Log,
	}
}

func (cs *CheckoutsService) GetDetailCheckout(token string) (*model.Checkouts, error) {

	checkout, err := cs.Repo.CheckoutRepo.GetDetailCheckout(token)
	if err != nil {
		cs.Logger.Error("Error from GetDetailCheckout :" + err.Error())
		return nil, err
	}

	productsArr := []model.Products{}

	for _, product := range *checkout.Products {
		product.PriceAfterDiscount = product.Price - (product.Price * product.Discount / 100)

		product.Amount = product.Qty * product.PriceAfterDiscount

		product.PriceAfterDiscount = 0
		product.Price = 0
		product.Qty = 0
		product.Discount = 0

		productsArr = append(productsArr, product)

	}

	checkouts := &model.Checkouts{
		Users:    checkout.Users, // Ambil data user dari repository
		Products: &productsArr,   // Produk yang telah diproses
	}

	return checkouts, nil
}
