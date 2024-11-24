package productshandler

import (
	"ecommers/helper"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (wh *ProductHandler) CreateWishlist(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	if err := wh.Service.ProductService.CreateWishlist(token, id); err != nil {
		wh.Log.Error("Failed to create wistlist: " + err.Error())
		helper.Responses(w, http.StatusInternalServerError, "Failed to create wishlist: "+err.Error(), nil)
		return
	}

	helper.Responses(w, http.StatusCreated, "Succsessfully", id)

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
