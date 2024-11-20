package products

import "ecommers/model"

func (c *ProductRepository) TotalCarts() (*[]model.Cart, error) {

	query := `SELECT u.id, COUNT(c.id) as total_products FROM shopping_carts c
		LEFT JOIN users u ON c.user_id = u.id
		WHERE c.deleted_at IS NULL AND c.user_id = 
		GROUP BY u.id`

	rows, err := c.DB.Query(query)
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
