package productshandler

import (
	"ecommers/helper"
	"net/http"
)

func (ch *ProductHandler) AllProductsCart(w http.ResponseWriter, r *http.Request) {

	carts, err := ch.Service.ProductService.TotalProducts()
	if err != nil {
		ch.Log.Error("Failed to Get total product cart: " + err.Error())
		helper.Responses(w, http.StatusInternalServerError, "Failed to Get total product cart: "+err.Error(), nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Succesfully", carts)
}
