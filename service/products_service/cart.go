package productsservice

import "ecommers/model"

func (cs *ProductsService) TotalProducts() (*[]model.Cart, error) {

	carts, err := cs.Repo.ProductsRepo.TotalCarts()
	if err != nil {
		cs.Logger.Error("Error from service: " + err.Error())
		return nil, err
	}

	return carts, nil
}
