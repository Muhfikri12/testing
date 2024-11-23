package checkoutservice

func (cs *CheckoutsService) CreateOrder(token string) error {

	err := cs.Repo.CheckoutRepo.ProcessCheckout(token)
	if err != nil {
		cs.Logger.Error("Error from service CreateOrder: " + err.Error())
		return err
	}

	return nil
}
