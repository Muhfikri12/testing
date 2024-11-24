package checkouthandler

import (
	"ecommers/helper"
	"net/http"
)

func (ch *CheckoutsHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")

	checkout, err := ch.Service.CheckoutService.CreateOrder(token)
	if err != nil {
		ch.Log.Error("Product not found: " + err.Error())
		helper.Responses(w, http.StatusNotFound, "Product not found: "+err.Error(), nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Succesfully", checkout)
}
