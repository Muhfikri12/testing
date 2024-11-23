package cartshandler

import (
	"ecommers/helper"
	"net/http"
)

func (ch *CartsHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")

	err := ch.Service.CartService.CreateOrder(token)
	if err != nil {
		ch.Log.Error("Product not found: " + err.Error())
		helper.Responses(w, http.StatusNotFound, "Product not found: "+err.Error(), nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Succesfully", nil)
}
