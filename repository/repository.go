package repository

import (
	"database/sql"
	"ecommers/repository/auth"
	"ecommers/repository/carts"
	"ecommers/repository/category"
	"ecommers/repository/checkout"
	"ecommers/repository/products"
	"ecommers/repository/promotions"
	"ecommers/repository/shipping"
	"ecommers/repository/users"

	"go.uber.org/zap"
)

type AllRepository struct {
	ProductsRepo  products.ProductRepository
	CategoryRepo  category.CategoryRepository
	PromotionRepo promotions.PromotionsRepository
	UsersRepo     users.UsersRepoInterface
	CartRepo      carts.CartsRepository
	CheckoutRepo  checkout.CheckoutsRepository
	AuthRepo      auth.AuthRepositoryInterface
	ShippingRepo  shipping.ShippingRepository
}

func NewAllRepository(db *sql.DB, log *zap.Logger) AllRepository {
	return AllRepository{
		ProductsRepo:  products.NewProductRepository(db, log),
		CategoryRepo:  category.NewCategoryRepository(db, log),
		PromotionRepo: promotions.NewPromotionsRepository(db, log),
		UsersRepo:     users.NewUsersRepository(db, log),
		CartRepo:      carts.NewCartsRepository(db, log),
		CheckoutRepo:  checkout.NewCheckoutsRepository(db, log),
		AuthRepo:      auth.NewAuthRepository(db, log),
		ShippingRepo:  shipping.NewShippingRepository(log, db),
	}
}
