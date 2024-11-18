package handler

import (
	productshandler "ecommers/handler/products"
	"ecommers/service"
	"ecommers/util"

	"go.uber.org/zap"
)

type AllHandler struct {
	SampelHandler  SampelHandler
	ProductHandler productshandler.ProductHandler
}

func NewAllHandler(service service.AllService, log *zap.Logger, config util.Configuration) AllHandler {
	return AllHandler{
		SampelHandler:  NewSampelService(service, log, config),
		ProductHandler: productshandler.NewProductsHandler(service, log, config),
	}

}
