package checkoutservice

import (
	"ecommers/model"
	"ecommers/repository"
	"errors"

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

func (cs *CheckoutsService) GetDetailCheckout(token, expedisi string) (*model.Checkouts, error) {

	checkout, err := cs.Repo.CheckoutRepo.GetDetailCheckout(token)
	if err != nil {
		cs.Logger.Error("Error from GetDetailCheckout: " + err.Error())
		return nil, err
	}

	distance, err := cs.Repo.ShippingRepo.ShippingCounting(checkout.Users.Address.Longlat)
	if err != nil {
		cs.Logger.Error("Failed Get distance: " + err.Error())
		return nil, err
	}

	distanceInKm := distance / 1000.0

	totalQty := 0
	for _, product := range *checkout.Products {
		totalQty += product.Qty
	}

	var costPerKm float64
	shippingCosts := map[string]map[bool]float64{
		"JNE":   {true: 2000, false: 4000}, // true: < 2 barang, false: >= 2 barang
		"JNT":   {true: 3000, false: 5000},
		"Ninja": {true: 3500, false: 5500},
	}

	// Validasi ekspedisi dan ambil biaya
	if costs, exists := shippingCosts[expedisi]; exists {
		costPerKm = costs[totalQty < 2]
	} else {
		cs.Logger.Error("Invalid expedisi: " + expedisi)
		return nil, errors.New("invalid expedisi")
	}
	shippingCost := costPerKm * distanceInKm

	var productsArr []model.Products
	totalAmount := 0
	for _, product := range *checkout.Products {

		product.PriceAfterDiscount = product.Price - (product.Price * product.Discount / 100)

		product.Amount = product.Qty * product.PriceAfterDiscount

		totalAmount += product.Amount

		product.PriceAfterDiscount = 0
		product.Price = 0
		product.Discount = 0

		productsArr = append(productsArr, product)
	}

	checkouts := &model.Checkouts{
		Users:        checkout.Users,
		ShippingCost: int(shippingCost),
		TotalAmount:  totalAmount + int(shippingCost),
		Products:     &productsArr,
	}

	return checkouts, nil
}
