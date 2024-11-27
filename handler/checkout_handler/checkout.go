package checkouthandler

import (
	"ecommers/helper"
	"ecommers/service"
	"ecommers/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
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

	expedisi := chi.URLParam(r, "expedisi")
	checkout, err := ch.Service.CheckoutService.GetDetailCheckout(token, expedisi)
	if err != nil {
		ch.Log.Error("Data is not found" + err.Error())
		helper.Responses(w, http.StatusNotFound, "Data Not Found", nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Successfully Get Data", checkout)
}

func (ch *CheckoutsHandler) GetDetailCheckoutGin(c *gin.Context) {
	// Ambil token dari header Authorization
	token := c.GetHeader("Authorization")
	if token == "" {
		ch.Log.Error("Authorization token is missing")
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	var expedisi string = c.Query("expedisi")
	if expedisi == "" {
		helper.ResponsesJson(c, http.StatusUnprocessableEntity, "Expedisi parameter is required", nil)
		return
	}

	// Panggil service untuk mendapatkan detail checkout
	checkout, err := ch.Service.CheckoutService.GetDetailCheckout(token, expedisi)
	if err != nil {
		ch.Log.Error("Data not found: " + err.Error())
		c.JSON(http.StatusNotFound, err)
		return
	}

	// Kembalikan response sukses
	helper.ResponsesJson(c, http.StatusOK, "success Get Data", checkout)
}
