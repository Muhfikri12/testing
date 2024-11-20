package service

import (
	"ecommers/repository"
	authservice "ecommers/service/auth_service"
	categoryservice "ecommers/service/category_service"
	productsservice "ecommers/service/products_service"
	promotionsservice "ecommers/service/promotions_service"

	"go.uber.org/zap"
)

type AllService struct {
	ProductService   productsservice.ProductsService
	CategoryService  categoryservice.CategoriesService
	PromotionService promotionsservice.PromotionsService
	AuthService      authservice.AuthService
}

func NewAllService(repo repository.AllRepository, log *zap.Logger) AllService {
	return AllService{
		ProductService:   productsservice.NewProductsService(repo, log),
		CategoryService:  categoryservice.NewCategoriesService(repo, log),
		PromotionService: promotionsservice.NewPromotionsService(repo, log),
		AuthService:      authservice.NewAuthService(repo, log),
	}
}
