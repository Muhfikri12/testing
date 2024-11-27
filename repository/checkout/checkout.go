package checkout

import (
	"database/sql"
	"ecommers/model"

	"go.uber.org/zap"
)

type CheckoutsRepository struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewCheckoutsRepository(db *sql.DB, Log *zap.Logger) CheckoutsRepository {
	return CheckoutsRepository{
		DB:     db,
		Logger: Log,
	}
}

func (c *CheckoutsRepository) GetDetailCheckout(token string) (*model.Checkouts, error) {

	User := model.Users{
		Address: &model.Addresses{},
	}
	Products := []model.Products{}

	queryUser := `SELECT u.name, u.email, a.address, a.longlat FROM shopping_carts c
			JOIN users u ON c.user_id = u.id
			JOIN addresses a ON u.id = a.user_id
			WHERE u.token = $1 AND a.is_main = true`

	err := c.DB.QueryRow(queryUser, token).Scan(&User.Name, &User.Email, &User.Address.Address, &User.Address.Longlat)
	if err != nil {
		c.Logger.Error("Error from query users GetDetailCheckout: " + err.Error())
		return nil, err
	}

	query := `SELECT p.name, p.image_url, p.price, p.discount, SUM(c.qty) FROM shopping_carts c
		JOIN product_varians pv ON c.product_variant_id = pv.id
		JOIN products p ON pv.product_id = p.id
		JOIN users u ON c.user_id = u.id
		JOIN addresses a ON u.id = a.user_id
		WHERE pv.product_id = p.id AND u.token = $1
		GROUP BY p.name, p.image_url, p.price, p.discount, u.name, u.email, a.address`

	rows, err := c.DB.Query(query, token)
	if err != nil {
		c.Logger.Error("Error from query GetDetailCheckout: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := model.Products{}
		if err := rows.Scan(&product.Name, &product.ImageUrl, &product.Price, &product.Discount, &product.Qty); err != nil {
			c.Logger.Error("Error from scan GetDetailCheckout: " + err.Error())
			return nil, err
		}

		Products = append(Products, product)
	}

	return &model.Checkouts{
		Users:    &User,
		Products: &Products,
	}, nil
}
