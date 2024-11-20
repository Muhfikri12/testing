package productshandler

import (
	"ecommers/helper"
	"ecommers/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func (wh *ProductHandler) CreateWishlist(w http.ResponseWriter, r *http.Request) {

	wishlist := model.Wishlists{}
	validate := validator.New()

	if err := json.NewDecoder(r.Body).Decode(&wishlist); err != nil {
		wh.Log.Error("Error decoding JSON: " + err.Error())
		helper.Responses(w, http.StatusBadRequest, "Invalid JSON payload", nil)
		return
	}

	err := validate.Struct(wishlist)
	if err != nil {
		errors, _ := helper.ValidateInputGeneric(wishlist)
		helper.Responses(w, http.StatusUnprocessableEntity, "validation failed", errors)
		return
	}

	if err := wh.Service.ProductService.CreateWishlist(&wishlist); err != nil {
		wh.Log.Error("Failed to create wistlist: " + err.Error())
		helper.Responses(w, http.StatusInternalServerError, "Failed to create wishlist: "+err.Error(), nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Succsessfully", wishlist)

}

func (wh *ProductHandler) DeleteWishlist(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		wh.Log.Error("event handler: id parameter missing")
		helper.Responses(w, http.StatusBadRequest, "id parameter is required", nil)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		wh.Log.Error("event handler: " + err.Error())
		helper.Responses(w, http.StatusBadRequest, "invalid id parameter", nil)
		return
	}

	if err := wh.Service.ProductService.DeleteWishlist(id); err != nil {
		wh.Log.Error("event handler:" + err.Error())
		helper.Responses(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Succsessfully", id)
}
