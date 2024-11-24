package carts

import (
	"database/sql"
	"ecommers/model"
	"errors"
	"time"

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

	query := `SELECT COUNT(c.id) as total_products FROM shopping_carts c
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
		WHERE pv.product_id = p.id AND u.token = $1 AND c.deleted_at IS NULL
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

func (c *CartsRepository) AddItemToCart(token string, cart *model.Products) error {

	userID, err := c.GetUserID(token)
	if err != nil {
		c.Logger.Error("Error For Getting id User: " + err.Error())
		return err
	}

	today := time.Now()

	queryAddToCart := `
		INSERT INTO shopping_carts (product_variant_id, user_id, qty, created_at, updated_at) 
		VALUES ($1, $2, 1, $3, $4)
		ON CONFLICT (product_variant_id, user_id) 
		DO UPDATE SET qty = shopping_carts.qty + 1, updated_at = $5
		WHERE shopping_carts.deleted_at IS NULL
	`
	_, err = c.DB.Exec(queryAddToCart, cart.ID, userID, today, today, today)
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

func (c *CartsRepository) UpdateCart(token string, product *model.Products) error {
	today := time.Now()

	userID, err := c.GetUserID(token)
	if err != nil {
		c.Logger.Error("Error For Getting User ID: " + err.Error())
		return err
	}

	query := `
		UPDATE shopping_carts 
		SET qty = $1, updated_at = $2
		WHERE product_variant_id = $3 AND user_id = $4 AND deleted_at IS NULL
	`
	result, err := c.DB.Exec(query, product.Qty, today, product.ID, userID)
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

func (c *CartsRepository) DeleteCart(token string, product *model.Products) error {
	today := time.Now()

	userID, err := c.GetUserID(token)
	if err != nil {
		c.Logger.Error("Error getting User ID: " + err.Error())
		return err
	}

	query := `
		UPDATE shopping_carts 
		SET deleted_at = $1
		WHERE deleted_at IS NULL AND product_variant_id = $2 AND user_id = $3
	`
	result, err := c.DB.Exec(query, today, product.ID, userID)
	if err != nil {
		c.Logger.Error("Error soft deleting cart item: " + err.Error())
		return err
	}

	// Periksa apakah ada baris yang diperbarui
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
