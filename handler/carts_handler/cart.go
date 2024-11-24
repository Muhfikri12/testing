package cartshandler

import (
	"ecommers/helper"
	"ecommers/model"
	"ecommers/service"
	"ecommers/util"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

	totalProduct, err := ch.Service.CartService.TotalProducts(token)
	if err != nil {
		ch.Log.Error("Failed to Get total product cart: " + err.Error())
		helper.Responses(w, http.StatusInternalServerError, "Failed to Get total product cart: "+err.Error(), nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Succesfully", totalProduct)
}

func (ch *CartsHandler) GetDetailCart(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")

	carts, err := ch.Service.CartService.GetDetailCart(token)
	if err != nil {
		ch.Log.Error("Product not found: " + err.Error())
		helper.Responses(w, http.StatusNotFound, "Product not found: "+err.Error(), nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Succesfully", carts)
}

func (ch *CartsHandler) AddItemToCart(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")
	idSrt := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idSrt)

	if err := ch.Service.CartService.AddItemToCart(token, id); err != nil {
		ch.Log.Error("Failed to Insert Product to cart: " + err.Error())
		helper.Responses(w, http.StatusInternalServerError, "failed to insert product to cart: "+err.Error(), nil)
		return
	}

	helper.Responses(w, http.StatusCreated, "Successfully Insert to cart", nil)
}

func (ch *CartsHandler) UpdateCart(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")
	idSrt := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idSrt)
	product := model.Products{}

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		ch.Log.Error("Error from Decode UpdateCart: " + err.Error())
		return
	}

	if err := ch.Service.CartService.UpdateCart(token, id, &product); err != nil {
		ch.Log.Error("Failed to Update Qty: " + err.Error())
		helper.Responses(w, http.StatusNotFound, "Product not found", nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Successfully Update qty to cart", product)
}

func (ch *CartsHandler) DeleteCart(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")
	product := model.Products{}

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		ch.Log.Error("Error from Decode DeleteCart: " + err.Error())
		return
	}

	if err := ch.Service.CartService.DeleteCart(token, &product); err != nil {
		ch.Log.Error("Failed to Delete item: " + err.Error())
		helper.Responses(w, http.StatusNotFound, "Product not found", nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Successfully deleting item", product)
}
