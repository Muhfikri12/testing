package promotionshandler

import (
	"ecommers/helper"
	"ecommers/service"
	"ecommers/util"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (bh *PromotionsHandler) GetAllBanners(c *gin.Context) {
	status := false
	days := 7

	banners, err := bh.Service.PromotionService.GetallCampaign(status, days)
	if err != nil {
		bh.Log.Error("Error fatch banners from handler: " + err.Error())
		helper.ResponsesJson(c, http.StatusInternalServerError, "Error Fatch banners"+err.Error(), nil)
		return
	}

	helper.ResponsesJson(c, http.StatusOK, "Successfully Fatch Data", banners)
}

func (ph *PromotionsHandler) GetAllPromo(c *gin.Context) {
	status := true
	days := 30

	banners, err := ph.Service.PromotionService.GetallCampaign(status, days)
	if err != nil {
		ph.Log.Error("Error fatch banners from handler: " + err.Error())
		helper.ResponsesJson(c, http.StatusInternalServerError, "Error Fatch banners"+err.Error(), nil)
		return
	}

	helper.ResponsesJson(c, http.StatusOK, "Successfully Fatch Data", banners)
}

func (rh *PromotionsHandler) GetAllRecomended(c *gin.Context) {

	recoments, err := rh.Service.PromotionService.GetAllRecomended()
	if err != nil {
		rh.Log.Error("Error fatch recoments from handler: " + err.Error())
		helper.ResponsesJson(c, http.StatusInternalServerError, "Error Fatch recoments"+err.Error(), nil)
		return
	}

	helper.ResponsesJson(c, http.StatusOK, "Successfully Fatch Data", recoments)

}
