package productsservice

import "ecommers/model"

func (cs *ProductsService) TotalProducts(token string) (*[]model.Cart, error) {

	carts, err := cs.Repo.ProductsRepo.TotalCarts(token)
	if err != nil {
		cs.Logger.Error("Error from service: " + err.Error())
		return nil, err
	}

	return carts, nil
}
