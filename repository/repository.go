package repository

import (
	"database/sql"
	"ecommers/repository/category"
	"ecommers/repository/products"

	"go.uber.org/zap"
)

type AllRepository struct {
	ProductsRepo products.ProductRepository
	CategoryRepo category.CategoryRepository
}

func NewAllRepository(db *sql.DB, log *zap.Logger) AllRepository {
	return AllRepository{
		ProductsRepo: products.NewProductRepository(db, log),
		CategoryRepo: category.NewCategoryRepository(db, log),
	}
}
