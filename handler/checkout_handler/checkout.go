package checkouthandler

import (
	"ecommers/helper"
	"ecommers/service"
	"ecommers/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CheckoutsHandler struct {
	Service service.AllService
	Log     *zap.Logger
	Config  util.Configuration
}

func NewCheckoutsHandler(service service.AllService, log *zap.Logger, config util.Configuration) CheckoutsHandler {
	return CheckoutsHandler{
		Service: service,
		Log:     log,
		Config:  config,
	}
}

func (ch *CheckoutsHandler) GetDetailCheckoutGin(c *gin.Context) {

	token := c.GetHeader("Authorization")

	var expedisi string = c.Query("expedisi")
	if expedisi == "" {
		helper.ResponsesJson(c, http.StatusUnprocessableEntity, "Expedisi parameter is required", nil)
		return
	}

	checkout, err := ch.Service.CheckoutService.GetDetailCheckout(token, expedisi)
	if err != nil {
		ch.Log.Error("Data not found: " + err.Error())
		helper.ResponsesJson(c, http.StatusNotFound, "Not item for check out", nil)
		return
	}

	helper.ResponsesJson(c, http.StatusOK, "success Get Data", checkout)
}
