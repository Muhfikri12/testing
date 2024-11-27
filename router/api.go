package router

import (
	"ecommers/database"
	"ecommers/handler"
	"ecommers/middleware"
	"ecommers/repository"
	"ecommers/service"
	"ecommers/util"
	"log"

	"github.com/gin-gonic/gin"
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
			r.Use(auth.AuthenticateToken)
			r.Post("/{id}", handler.ProductHandler.CreateWishlist)
			r.Delete("/{id}", handler.ProductHandler.DeleteWishlist)
		})

		api.Route("/carts", func(r chi.Router) {
			r.Use(auth.AuthenticateToken)
			r.Get("/", handler.CartHandler.AllProductsCart)
			// r.Post("/{id}", handler.CartHandler.AddItemToCart)
			r.Get("/detail", handler.CartHandler.GetDetailCart)
			r.Put("/{id}", handler.CartHandler.UpdateCart)
			r.Delete("/{id}", handler.CartHandler.DeleteCart)

		})

		api.Route("/users", func(r chi.Router) {
			r.Use(auth.AuthenticateToken)
			r.Get("/", handler.UserHandler.GetListAddress)
			r.Put("/", handler.UserHandler.UpdateUserData)
			r.Get("/detail", handler.UserHandler.GetDetailUser)
			r.Post("/address", handler.UserHandler.AddAddress)
			r.Put("/address/{id}", handler.UserHandler.UpdateAddress)
			r.Delete("/address/{id}", handler.UserHandler.DeleteAddress)
			r.Patch("/address/set/{id}", handler.UserHandler.SetDefault)
		})

		api.Route("/order", func(r chi.Router) {
			r.Use(auth.AuthenticateToken)
			r.Post("/", handler.Checkouthandler.CreateOrder)
		})

		api.Post("/login", handler.AuthHandler.Login)
		api.Post("/register", handler.AuthHandler.Register)
		api.Get("/categories", handler.CategoryHandler.GetAllCategories)
		api.Get("/product/{id}", handler.ProductHandler.GetProductByID)
		api.Get("/banners", handler.PromotionHandler.GetAllBanners)
		api.Get("/promo", handler.PromotionHandler.GetAllPromo)
		api.Get("/recomended", handler.PromotionHandler.GetAllRecomended)

	})

	return r, logger, nil
}

// InitGin initializes the Gin application
func InitGin() (*gin.Engine, *zap.Logger, error) {
	// Initialize Logger
	logger := util.InitLog()

	// Load Configuration
	config, err := util.ReadConfiguration()
	if err != nil {
		logger.Error("Failed to read configuration", zap.Error(err))
		return nil, logger, err
	}

	// Initialize Database
	db, err := database.InitDB(config)
	if err != nil {
		logger.Error("Failed to initialize database", zap.Error(err))
		return nil, logger, err
	}

	// Initialize Authentication Middleware
	auth := middleware.NewAuthHandler(logger)

	// Initialize Repository, Service, and Handlers
	repo := repository.NewAllRepository(db, logger)
	service := service.NewAllService(repo, logger)
	allHandler := handler.NewAllHandler(service, logger, config)

	// Create Gin Router
	router := gin.Default()

	router.POST("/login", allHandler.AuthHandler.LoginGin)

	authorized := router.Group("/")
	authorized.Use(auth.AuthenticateGin())
	{
		authorized.GET("/checkouts", allHandler.Checkouthandler.GetDetailCheckoutGin)
	}

	cart := router.Group("/carts")
	cart.Use(auth.AuthenticateGin())
	{
		cart.POST("/:id", allHandler.CartHandler.AddItemToCart)
	}

	// Return the configured router and logger
	log.Println("Gin application initialized")
	return router, logger, nil
}
