package cartshandler

import (
	"ecommers/helper"
	"ecommers/service"
	"ecommers/util"
	"net/http"

	"go.uber.org/zap"
)

type CartsHandler struct {
	Service service.AllService
	Log     *zap.Logger
	Config  util.Configuration
}

func NewCartssHandler(service service.AllService, log *zap.Logger, config util.Configuration) CartsHandler {
	return CartsHandler{
		Service: service,
		Log:     log,
		Config:  config,
	}
}

func (ch *CartsHandler) AllProductsCart(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")

	carts, err := ch.Service.CartService.TotalProducts(token)
	if err != nil {
		ch.Log.Error("Failed to Get total product cart: " + err.Error())
		helper.Responses(w, http.StatusInternalServerError, "Failed to Get total product cart: "+err.Error(), nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Succesfully", carts)
}
