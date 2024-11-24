package productsservice

func (ws *ProductsService) CreateWishlist(token string, id int) error {

	err := ws.Repo.ProductsRepo.CreateWishlisht(token, id)
	if err != nil {
		ws.Logger.Error("Error from Service: " + err.Error())
		return err
	}

	return nil
}

func (ws *ProductsService) DeleteWishlist(id int, token string) error {

	err := ws.Repo.ProductsRepo.DeleteWishlist(id, token)
	if err != nil {
		ws.Logger.Error("Error from Service: " + err.Error())
		return err
	}

	return nil
}
