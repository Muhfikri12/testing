package repository

import (
	"database/sql"
	"ecommers/repository/category"
	"ecommers/repository/products"
	"ecommers/repository/promotions"
	"ecommers/repository/users"

	"go.uber.org/zap"
)

type AllRepository struct {
	ProductsRepo  products.ProductRepository
	CategoryRepo  category.CategoryRepository
	PromotionRepo promotions.PromotionsRepository
	UsersRepo     users.UsersRepository
}

func NewAllRepository(db *sql.DB, log *zap.Logger) AllRepository {
	return AllRepository{
		ProductsRepo:  products.NewProductRepository(db, log),
		CategoryRepo:  category.NewCategoryRepository(db, log),
		PromotionRepo: promotions.NewPromotionsRepository(db, log),
		UsersRepo:     users.NewUsersRepository(db, log),
	}
}
