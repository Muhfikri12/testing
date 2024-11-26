package carts

import (
	"database/sql"
	"ecommers/model"
	"errors"

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

func (c *CartsRepository) TotalCarts(token string) (int, error) {

	var totalProduct int
	query := `SELECT SUM(qty) as total_products FROM shopping_carts 
		WHERE user_id = (SELECT id FROM users WHERE token=$1)`

	err := c.DB.QueryRow(query, token).Scan(&totalProduct)
	if err != nil {
		c.Logger.Error("Error from repository: " + err.Error())
		return 0, err
	}

	return totalProduct, nil
}

func (c *CartsRepository) GetDetailCart(token string) (*[]model.Products, error) {

	carts := []model.Products{}

	query := `SELECT p.name, pv.image_url, p.price, p.discount, SUM(c.qty) FROM shopping_carts c
		JOIN product_varians pv ON c.product_variant_id = pv.id
		JOIN products p ON pv.product_id = p.id
		JOIN users u ON c.user_id = u.id
		WHERE pv.product_id = p.id AND u.token = $1 AND c.deleted_at IS NULL
		GROUP BY p.name, pv.image_url, p.price, p.discount`

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

func (c *CartsRepository) AddItemToCart(token string, id int) error {

	queryAddToCart := `
		INSERT INTO shopping_carts (product_variant_id, user_id, qty, created_at, updated_at) 
		VALUES ($1, (SELECT id FROM users WHERE token=$2), 1, NOW(), NOW())
		ON CONFLICT (product_variant_id, user_id) 
		DO UPDATE SET qty = shopping_carts.qty + 1, updated_at = NOW()
		WHERE shopping_carts.deleted_at IS NULL
	`
	_, err := c.DB.Exec(queryAddToCart, id, token)
	if err != nil {
		c.Logger.Error("Error For input or update database: " + err.Error())
		return err
	}

	return nil
}

func (c *CartsRepository) GetUserID(token string) (int, error) {

	var userID int
	queryGetUserID := `
		SELECT id
		FROM users
		WHERE token = $1
	`
	err := c.DB.QueryRow(queryGetUserID, token).Scan(&userID)
	if err != nil {
		c.Logger.Error("Error from Get Id User: " + err.Error())
		return 0, err
	}

	return userID, err
}

func (c *CartsRepository) UpdateCart(token string, id int, product *model.Products) error {

	query := `
		UPDATE shopping_carts 
		SET qty = $1, updated_at = NOW()
		WHERE product_variant_id = $2 AND user_id = (SELECT id FROM users WHERE token=$3) AND deleted_at IS NULL
	`
	result, err := c.DB.Exec(query, product.Qty, id, token)
	if err != nil {
		c.Logger.Error("Error updating cart: " + err.Error())
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.Logger.Error("Error getting RowsAffected: " + err.Error())
		return err
	}

	if rowsAffected == 0 {
		c.Logger.Warn("No cart item deleted, possibly already soft deleted or not found")
		return errors.New("item not found or already deleted")
	}

	return nil
}

func (c *CartsRepository) DeleteCart(token string, id int) error {
	query := `
		DELETE FROM shopping_carts
		WHERE product_variant_id = $1 AND user_id = (SELECT id FROM users WHERE token=$2)
	`
	result, err := c.DB.Exec(query, id, token)
	if err != nil {
		c.Logger.Error("Error soft deleting cart item: " + err.Error())
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.Logger.Error("Error getting RowsAffected: " + err.Error())
		return err
	}

	if rowsAffected == 0 {
		c.Logger.Warn("No cart item deleted, possibly already soft deleted or not found")
		return errors.New("item not found or already deleted")
	}

	return nil
}
