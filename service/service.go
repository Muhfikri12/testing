package service

import (
	"ecommers/repository"
	categoryservice "ecommers/service/category_service"
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
}

func NewAllService(repo repository.AllRepository, log *zap.Logger) AllService {
	return AllService{
		ProductService:   productsservice.NewProductsService(repo, log),
		CategoryService:  categoryservice.NewCategoriesService(repo, log),
		PromotionService: promotionsservice.NewPromotionsService(repo, log),
		UserService:      usersservice.NewUsersService(repo, log),
	}
}
