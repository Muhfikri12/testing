package checkouthandler

import (
	"ecommers/helper"
	"ecommers/service"
	"ecommers/util"
	"net/http"

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

func (ch *CheckoutsHandler) GetDetailCheckout(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")

	checkout, err := ch.Service.CheckoutService.GetDetailCheckout(token)
	if err != nil {
		ch.Log.Error("Data is not found" + err.Error())
		helper.Responses(w, http.StatusNotFound, "Data Not Found", nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Successfully Get Data", checkout)
}
