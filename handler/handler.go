package handler

import (
	categorieshandler "ecommers/handler/categories_handler"
	productshandler "ecommers/handler/products"
	promotionshandler "ecommers/handler/promotions_handler"
	"ecommers/service"
	"ecommers/util"

	"go.uber.org/zap"
)

type AllHandler struct {
	SampelHandler    SampelHandler
	ProductHandler   productshandler.ProductHandler
	CategoryHandler  categorieshandler.CategoriesHandler
	PromotionHandler promotionshandler.PromotionsHandler
}

func NewAllHandler(service service.AllService, log *zap.Logger, config util.Configuration) AllHandler {
	return AllHandler{
		SampelHandler:    NewSampelService(service, log, config),
		ProductHandler:   productshandler.NewProductsHandler(service, log, config),
		CategoryHandler:  categorieshandler.NewCategoriesHandler(service, log, config),
		PromotionHandler: promotionshandler.NewPromotionsHandler(service, log, config),
	}

}
