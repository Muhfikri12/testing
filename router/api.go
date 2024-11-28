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
	"go.uber.org/zap"
)

func InitGin() (*gin.Engine, *zap.Logger, error) {

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

	auth := middleware.NewAuthHandler(logger)

	repo := repository.NewAllRepository(db, logger)
	service := service.NewAllService(repo, logger)
	allHandler := handler.NewAllHandler(service, logger, config)

	md := middleware.NewMiddleware(logger)

	router := gin.Default()
	router.Use(md.MiddlewareLogger())
	middleware := router.Group("/api")

	middleware.POST("/login", allHandler.AuthHandler.LoginGin)
	middleware.POST("/register", allHandler.AuthHandler.Register)
	middleware.GET("/categories", allHandler.CategoryHandler.GetAllCategories)
	middleware.GET("/product/{id}", allHandler.ProductHandler.GetProductByID)
	middleware.GET("/banners", allHandler.PromotionHandler.GetAllBanners)
	middleware.GET("/promo", allHandler.PromotionHandler.GetAllPromo)
	middleware.GET("/recomended", allHandler.PromotionHandler.GetAllRecomended)

	middleware.Use(auth.AuthenticateGin())
	{
		middleware.GET("/checkouts", allHandler.Checkouthandler.GetDetailCheckoutGin)
	}

	cart := middleware.Group("/carts")
	{
		cart.POST("/:id", allHandler.CartHandler.AddItemToCart)
		cart.GET("/", allHandler.CartHandler.AllProductsCart)
		cart.GET("/detail", allHandler.CartHandler.GetDetailCart)
		cart.PUT("/:id", allHandler.CartHandler.UpdateCart)
		cart.DELETE("/:id:", allHandler.CartHandler.DeleteCart)
	}

	order := middleware.Group("/order")
	{
		order.POST("/", allHandler.Checkouthandler.CreateOrder)
	}

	wishlists := middleware.Group("/wishlist")
	{
		wishlists.POST("/:id", allHandler.ProductHandler.CreateWishlist)
		wishlists.DELETE("/:id", allHandler.ProductHandler.DeleteWishlist)
	}

	product := middleware.Group("/product")
	{
		product.GET("/", allHandler.ProductHandler.GetAll)
		product.GET("/best_selling", allHandler.ProductHandler.GetAllBestSelling)
	}

	users := middleware.Group("/user")
	{
		users.GET("/", allHandler.UserHandler.GetListAddress)
		users.PUT("/", allHandler.UserHandler.UpdateUserData)
		users.GET("/detail", allHandler.UserHandler.GetDetailUser)
		users.POST("/address", allHandler.UserHandler.AddAddress)
		users.PUT("/address/:id", allHandler.UserHandler.UpdateAddress)
		users.DELETE("/address/:id", allHandler.UserHandler.DeleteAddress)
		users.PATCH("/address/set/:id", allHandler.UserHandler.SetDefault)
	}

	// Return the configured router and logger
	log.Println("Gin application initialized")
	return router, logger, nil
}
