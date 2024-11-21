package handler

import (
	cartshandler "ecommers/handler/carts_handler"
	categorieshandler "ecommers/handler/categories_handler"
	checkouthandler "ecommers/handler/checkout_handler"
	productshandler "ecommers/handler/products"
	promotionshandler "ecommers/handler/promotions_handler"
	userhandler "ecommers/handler/user_handler"
	"ecommers/middleware"
	"ecommers/service"
	"ecommers/util"

	"go.uber.org/zap"
)

type AllHandler struct {
	ProductHandler   productshandler.ProductHandler
	CategoryHandler  categorieshandler.CategoriesHandler
	PromotionHandler promotionshandler.PromotionsHandler
	UserHandler      userhandler.UserHandler
	CartHandler      cartshandler.CartsHandler
	Checkouthandler  checkouthandler.CheckoutsHandler
	AuthHandler      middleware.AuthHandler
}

func NewAllHandler(service service.AllService, log *zap.Logger, config util.Configuration) AllHandler {
	return AllHandler{
		ProductHandler:   productshandler.NewProductsHandler(service, log, config),
		CategoryHandler:  categorieshandler.NewCategoriesHandler(service, log, config),
		PromotionHandler: promotionshandler.NewPromotionsHandler(service, log, config),
		UserHandler:      userhandler.NewUserHandler(service, log, config),
		CartHandler:      cartshandler.NewCartssHandler(service, log, config),
		Checkouthandler:  checkouthandler.NewCheckoutsHandler(service, log, config),
		AuthHandler:      middleware.NewAuthHandler(log),
	}

}
