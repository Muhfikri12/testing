package productsservice

import (
	"ecommers/model"
	"ecommers/repository"
	"fmt"
	"math"
	"sort"
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

	productsarr := []model.Products{}

	for _, product := range *products {

		if product.Timestamps.Created_at.Before(thirtyDaysAgo) {
			product.Status = "New"

		}

		product.PriceAfterDiscount = product.Price - (product.Price * product.Discount / 100)
		productsarr = append(productsarr, product)

	}

	return &productsarr, totalData, totalPage, nil
}

func (ps *ProductsService) ProductsBestSelling(page int, category, name string) (*[]model.Products, int, int, error) {
	limit := 10

	// Hitung awal dan akhir bulan ini
	thirtyDaysAgo := time.Now().AddDate(0, 0, 30)

	// Ambil data produk dari repository
	products, totalData, err := ps.Repo.ProductsRepo.ShowAllProducts(limit, page, category, name)
	if err != nil {
		ps.Logger.Error("Error fetching products: " + err.Error())
		return nil, 0, 0, err
	}

	if products == nil || len(*products) == 0 {
		ps.Logger.Warn("No products found")
		return &[]model.Products{}, 0, 0, nil
	}

	productsarr := []model.Products{}

	for _, product := range *products {

		if product.Timestamps.Created_at.Before(thirtyDaysAgo) {
			product.Status = "New"
		}

		if product.Timestamps.Created_at.Before(thirtyDaysAgo) {
			sort.Slice(productsarr, func(i, j int) bool {
				return productsarr[i].Checkouts.TotalSold > productsarr[j].Checkouts.TotalSold
			})
		}

		product.PriceAfterDiscount = product.Price - (product.Price * product.Discount / 100)
		productsarr = append(productsarr, product)

	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))

	return &productsarr, totalData, totalPage, nil
}
