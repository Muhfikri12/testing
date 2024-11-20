package promotionshandler

import (
	"ecommers/helper"
	"ecommers/service"
	"ecommers/util"
	"net/http"

	"go.uber.org/zap"
)

type PromotionsHandler struct {
	Service service.AllService
	Log     *zap.Logger
	Config  util.Configuration
}

func NewPromotionsHandler(service service.AllService, log *zap.Logger, config util.Configuration) PromotionsHandler {
	return PromotionsHandler{
		Service: service,
		Log:     log,
		Config:  config,
	}
}

func (bh *PromotionsHandler) GetAllBanners(w http.ResponseWriter, r *http.Request) {

	banners, err := bh.Service.PromotionService.GetallBanners()
	if err != nil {
		bh.Log.Error("Error fatch banners from handler: " + err.Error())
		helper.Responses(w, http.StatusInternalServerError, "Error Fatch banners"+err.Error(), nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Successfully Fatch Data", banners)
}
