package repository

import (
	"database/sql"
	"ecommers/repository/products"

	"go.uber.org/zap"
)

type AllRepository struct {
	SampelRepo   SampelRepository
	ProductsRepo products.ProductRepository
}

func NewAllRepository(db *sql.DB, log *zap.Logger) AllRepository {
	return AllRepository{
		SampelRepo:   NewSampelRepository(db, log),
		ProductsRepo: products.NewProductRepository(db, log),
	}
}
