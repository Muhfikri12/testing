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
	id, _ := strconv.Atoi(idStr)

	token := r.Header.Get("Authorization")

	if err := wh.Service.ProductService.DeleteWishlist(id, token); err != nil {
		wh.Log.Error("event handler:" + err.Error())
		helper.Responses(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Succsessfully", id)
}
