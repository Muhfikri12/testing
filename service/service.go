package service

import (
	"ecommers/repository"
	productsservice "ecommers/service/products_service"

	"go.uber.org/zap"
)

type AllService struct {
	ProductService productsservice.ProductsService
}

func NewAllService(repo repository.AllRepository, log *zap.Logger) AllService {
	return AllService{
		ProductService: productsservice.NewProductsService(repo, log),
	}
}
