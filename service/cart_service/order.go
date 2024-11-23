package cartservice

func (cs *CartsService) CreateOrder(token string) error {

	err := cs.Repo.CartRepo.ProcessCheckout(token)
	if err != nil {
		cs.Logger.Error("Error from service CreateOrder: " + err.Error())
		return err
	}

	return nil
}
