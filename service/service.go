package service

import (
	"ecommers/repository"
	categoryservice "ecommers/service/category_service"
	productsservice "ecommers/service/products_service"

	"go.uber.org/zap"
)

type AllService struct {
	ProductService  productsservice.ProductsService
	CategoryService categoryservice.CategoriesService
}

func NewAllService(repo repository.AllRepository, log *zap.Logger) AllService {
	return AllService{
		ProductService:  productsservice.NewProductsService(repo, log),
		CategoryService: categoryservice.NewCategoriesService(repo, log),
	}
}
