package productsservice

import (
	"ecommers/model"
	"ecommers/repository"
	"fmt"
	"math"
	"time"

	"go.uber.org/zap"
)

type ProductsService struct {
	Repo   repository.AllRepository
	Logger *zap.Logger
}

func NewProductsService(repo repository.AllRepository, Log *zap.Logger) ProductsService {
	return ProductsService{
		Repo:   repo,
		Logger: Log,
	}
}

func (ps *ProductsService) GetAll(page int, category, name string) (*[]model.Products, int, int, error) {
	limit := 10

	thirtyDaysAgo := time.Now().AddDate(0, 0, 30)

	if category == "" {
		category = ""
	}
	if name == "" {
		name = ""
	}

	products, totalData, err := ps.Repo.ProductsRepo.ShowAllProducts(limit, page, category, name)
	if err != nil {
		ps.Logger.Error("Error from Service: " + err.Error())
		return nil, 0, 0, err
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))

	if products == nil {
		ps.Logger.Error("No products found")
		return nil, 0, 0, fmt.Errorf("no products found")
	}

	productsarr := make([]model.Products, len(*products))
	copy(productsarr, *products)

	for i := range productsarr {
		product := &productsarr[i]

		if product.Timestamps.Created_at.Before(thirtyDaysAgo) {
			product.Status = "New"
		}
	}

	return &productsarr, totalData, totalPage, nil
}
