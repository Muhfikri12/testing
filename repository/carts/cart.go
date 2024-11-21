package carts

import (
	"database/sql"
	"ecommers/model"

	"go.uber.org/zap"
)

type CartsRepository struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewCartsRepository(db *sql.DB, Log *zap.Logger) CartsRepository {
	return CartsRepository{
		DB:     db,
		Logger: Log,
	}
}

func (c *CartsRepository) TotalCarts(token string) (*[]model.Cart, error) {

	query := `SELECT u.id, COUNT(c.id) as total_products FROM shopping_carts c
		LEFT JOIN users u ON c.user_id = u.id
		WHERE c.deleted_at IS NULL AND u.token = $1
		GROUP BY u.id`

	rows, err := c.DB.Query(query, token)
	if err != nil {
		c.Logger.Error("Error from repository: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	carts := []model.Cart{}

	for rows.Next() {
		cart := model.Cart{}
		if err := rows.Scan(&cart.UserID, &cart.TotalProducts); err != nil {
			c.Logger.Error("Error from repository: " + err.Error())
			return nil, err
		}

		carts = append(carts, cart)
	}

	return &carts, nil
}

func (c *CartsRepository) GetDetailCart(token string) (*[]model.Products, error) {

	carts := []model.Products{}

	query := `SELECT p.name, p.image_url, p.price, p.discount, SUM(c.qty) FROM shopping_carts c
		JOIN product_varians pv ON c.product_variant_id = pv.id
		JOIN products p ON pv.product_id = p.id
		JOIN users u ON c.user_id = u.id
		WHERE pv.product_id = p.id AND u.token = $1
		GROUP BY p.name, p.image_url, p.price, p.discount`

	rows, err := c.DB.Query(query, token)
	if err != nil {
		c.Logger.Error("Error from query GetDetailCart: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		cart := model.Products{}
		if err := rows.Scan(&cart.Name, &cart.ImageUrl, &cart.Price, &cart.Discount, &cart.Qty); err != nil {
			c.Logger.Error("Error from scan GetDetailCart: " + err.Error())
			return nil, err
		}
		cart.PriceAfterDiscount = cart.Price - (cart.Price * cart.Discount / 100)
		cart.Price = 0
		cart.Discount = 0

		carts = append(carts, cart)
	}

	return &carts, nil
}
