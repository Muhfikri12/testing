package productsservice

import (
	"ecommers/model"
)

func (ws *ProductsService) CreateWishlist(wishlist *model.Wishlists) error {

	err := ws.Repo.ProductsRepo.CreateWishlisht(wishlist)
	if err != nil {
		ws.Logger.Error("Error from Service: " + err.Error())
		return err
	}

	return nil
}

func (ws *ProductsService) DeleteWishlist(id int) error {

	err := ws.Repo.ProductsRepo.DeleteWishlist(id)
	if err != nil {
		ws.Logger.Error("Error from Service: " + err.Error())
		return err
	}

	return nil
}
