package productshandler

import (
	"ecommers/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (wh *ProductHandler) CreateWishlist(c *gin.Context) {

	token := c.GetHeader("Authorization")
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := wh.Service.ProductService.CreateWishlist(token, id); err != nil {
		wh.Log.Error("Failed to create wistlist: " + err.Error())
		helper.ResponsesJson(c, http.StatusInternalServerError, "Failed to create wishlist: "+err.Error(), nil)
		return
	}

	helper.ResponsesJson(c, http.StatusCreated, "Succsessfully", id)

}

func (wh *ProductHandler) DeleteWishlist(c *gin.Context) {

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	token := c.GetHeader("Authorization")

	if err := wh.Service.ProductService.DeleteWishlist(id, token); err != nil {
		wh.Log.Error("event handler:" + err.Error())
		helper.ResponsesJson(c, http.StatusNotFound, err.Error(), nil)
		return
	}

	helper.ResponsesJson(c, http.StatusOK, "Succsessfully", id)
}
