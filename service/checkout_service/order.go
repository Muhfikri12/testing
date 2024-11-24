package checkoutservice

import "ecommers/model"

func (cs *CheckoutsService) CreateOrder(token string) (*model.Checkouts, error) {

	checkout, err := cs.Repo.CheckoutRepo.ProcessCheckout(token)
	if err != nil {
		cs.Logger.Error("Error from service CreateOrder: " + err.Error())
		return nil, err
	}

	return checkout, nil
}
