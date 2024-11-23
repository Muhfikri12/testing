package service

import (
	"ecommers/repository"
	authservice "ecommers/service/auth_service"
	cartservice "ecommers/service/cart_service"
	categoryservice "ecommers/service/category_service"
	checkoutservice "ecommers/service/checkout_service"
	productsservice "ecommers/service/products_service"
	promotionsservice "ecommers/service/promotions_service"
	usersservice "ecommers/service/users_service"

	"go.uber.org/zap"
)

type AllService struct {
	ProductService   productsservice.ProductsService
	CategoryService  categoryservice.CategoriesService
	PromotionService promotionsservice.PromotionsService
	UserService      usersservice.UsersService
	CartService      cartservice.CartsService
	CheckoutService  checkoutservice.CheckoutsService
	AuthService      authservice.AuthService
}

func NewAllService(repo repository.AllRepository, log *zap.Logger) AllService {
	return AllService{
		ProductService:   productsservice.NewProductsService(repo, log),
		CategoryService:  categoryservice.NewCategoriesService(repo, log),
		PromotionService: promotionsservice.NewPromotionsService(repo, log),
		UserService:      usersservice.NewUsersService(repo, log),
		CartService:      cartservice.NewCartsService(repo, log),
		CheckoutService:  checkoutservice.NewCheckoutsService(repo, log),
		AuthService:      authservice.NewAuthService(repo, log),
	}
}
