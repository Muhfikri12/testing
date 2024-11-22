package router

import (
	"ecommers/database"
	"ecommers/handler"
	"ecommers/middleware"
	"ecommers/repository"
	"ecommers/service"
	"ecommers/util"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func InitRouter() (*chi.Mux, *zap.Logger, error) {
	r := chi.NewRouter()

	logger := util.InitLog()

	config, err := util.ReadConfiguration()
	if err != nil {
		logger.Error("Failed to read configuration", zap.Error(err))
		return nil, logger, err
	}

	db, err := database.InitDB(config)
	if err != nil {
		logger.Error("Failed to initialize database", zap.Error(err))
		return nil, logger, err
	}

	md := middleware.NewMiddleware(logger)

	// Authentication
	auth := middleware.NewAuthHandler(logger)

	repo := repository.NewAllRepository(db, logger)
	service := service.NewAllService(repo, logger)
	handler := handler.NewAllHandler(service, logger, config)

	r.Route("/api", func(api chi.Router) {

		api.Use(md.MinddlewareLogger)
		api.Route("/products", func(r chi.Router) {
			r.Get("/", handler.ProductHandler.GetAll)
			r.Get("/best_selling", handler.ProductHandler.GetAllBestSelling)
		})

		api.Route("/wishlists", func(r chi.Router) {
			r.Use(handler.AuthHandler.AuthenticateToken)
			r.Post("/", handler.ProductHandler.CreateWishlist)
			r.Delete("/{id}", handler.ProductHandler.DeleteWishlist)
		})

		api.Route("/carts", func(r chi.Router) {
			r.Use(auth.AuthenticateToken)
			r.Get("/", handler.CartHandler.AllProductsCart)
			r.Post("/", handler.CartHandler.AddItemToCart)
			r.Get("/detail", handler.CartHandler.GetDetailCart)
			r.Put("/update-qty", handler.CartHandler.UpdateCart)
			r.Delete("/", handler.CartHandler.DeleteCart)

		})

		api.Route("/users", func(r chi.Router) {
			r.Use(auth.AuthenticateToken)
			r.Get("/", handler.UserHandler.GetListAddress)
		})

		api.Route("/checkout", func(r chi.Router) {
			r.Use(auth.AuthenticateToken)
			r.Get("/", handler.Checkouthandler.GetDetailCheckout)
		})

		api.Post("/login", handler.UserHandler.Login)
		api.Post("/register", handler.UserHandler.Register)
		api.Get("/categories", handler.CategoryHandler.GetAllCategories)
		api.Get("/product/{id}", handler.ProductHandler.GetProductByID)
		api.Get("/banners", handler.PromotionHandler.GetAllBanners)
		api.Get("/promo", handler.PromotionHandler.GetAllPromo)
		api.Get("/recomended", handler.PromotionHandler.GetAllRecomended)

	})

	return r, logger, nil
}
