package checkouthandler

import (
	"ecommers/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ch *CheckoutsHandler) CreateOrder(c *gin.Context) {

	token := c.GetHeader("Authorization")

	checkout, err := ch.Service.CheckoutService.CreateOrder(token)
	if err != nil {
		ch.Log.Error("Product not found: " + err.Error())
		helper.ResponsesJson(c, http.StatusNotFound, "Product not found", nil)
		return
	}

	helper.ResponsesJson(c, http.StatusOK, "Data retrieved successfully", checkout)
}
