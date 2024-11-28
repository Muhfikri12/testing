package cartshandler

import (
	"ecommers/helper"
	"ecommers/model"
	"ecommers/service"
	"ecommers/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (ch *CartsHandler) AllProductsCart(c *gin.Context) {

	token := c.GetHeader("Authorization")

	totalProduct, err := ch.Service.CartService.TotalProducts(token)
	if err != nil {
		ch.Log.Error("Failed to Get total product cart: " + err.Error())
		helper.ResponsesJson(c, http.StatusInternalServerError, "Failed to Get total product cart: "+err.Error(), nil)
		return
	}

	helper.ResponsesJson(c, http.StatusOK, "Succesfully", totalProduct)
}

func (ch *CartsHandler) GetDetailCart(c *gin.Context) {

	token := c.GetHeader("Authorization")

	carts, err := ch.Service.CartService.GetDetailCart(token)
	if err != nil {
		ch.Log.Error("Product not found: " + err.Error())
		helper.ResponsesJson(c, http.StatusNotFound, "Product not found: "+err.Error(), nil)
		return
	}

	helper.ResponsesJson(c, http.StatusOK, "Succesfully", carts)

}

func (ch *CartsHandler) AddItemToCart(c *gin.Context) {

	token := c.GetHeader("Authorization")
	idSrt := c.Param("id")
	id, _ := strconv.Atoi(idSrt)

	if err := ch.Service.CartService.AddItemToCart(token, id); err != nil {
		ch.Log.Error("Failed to Insert Product to cart: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to add item to cart",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Added item",
		"data":    id,
	})
}

func (ch *CartsHandler) UpdateCart(c *gin.Context) {

	token := c.GetHeader("Authorization")
	idSrt := c.Param("id")
	id, _ := strconv.Atoi(idSrt)
	product := model.Products{}

	err := c.ShouldBindJSON(&product)
	if err != nil {
		ch.Log.Error("Error from Decode UpdateCart: " + err.Error())
		return
	}

	if err := ch.Service.CartService.UpdateCart(token, id, &product); err != nil {
		ch.Log.Error("Failed to Update Qty: " + err.Error())
		helper.ResponsesJson(c, http.StatusNotFound, "Product not found", nil)
		return
	}

	helper.ResponsesJson(c, http.StatusOK, "Successfully Update qty to cart", product)

}

func (ch *CartsHandler) DeleteCart(c *gin.Context) {

	token := c.GetHeader("Authorization")

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := ch.Service.CartService.DeleteCart(token, id); err != nil {
		ch.Log.Error("Failed to Delete item: " + err.Error())
		helper.ResponsesJson(c, http.StatusNotFound, "Product not found", nil)
		return
	}

	helper.ResponsesJson(c, http.StatusOK, "Successfully deleting item", nil)
}
