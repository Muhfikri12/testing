package checkout

func (c *CheckoutsRepository) ProcessCheckout(token string) error {

	var userID int

	queryGetUserID := `
		SELECT id
		FROM users
		WHERE token = $1
	`
	err := c.DB.QueryRow(queryGetUserID, token).Scan(&userID)
	if err != nil {
		c.Logger.Error("Error from Get Id User: " + err.Error())
		return err
	}

	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	queryCartItems := `
        SELECT product_variant_id, qty
        FROM shopping_carts
        WHERE user_id = $1 AND deleted_at IS NULL
    `
	rows, err := tx.Query(queryCartItems, userID)
	if err != nil {
		return err
	}
	defer rows.Close()

	var cartItems []struct {
		ProductVariantID int
		Qty              int
	}

	for rows.Next() {
		var item struct {
			ProductVariantID int
			Qty              int
		}
		if err := rows.Scan(&item.ProductVariantID, &item.Qty); err != nil {
			return err
		}
		cartItems = append(cartItems, item)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	var totalAmount int
	for _, item := range cartItems {
		var price int

		queryGetPrice := `
            SELECT p.price - (p.price * p.discount / 100) as price
            FROM product_varians pv
            JOIN products p ON pv.product_id = p.id
            WHERE pv.id = $1
        `
		err = tx.QueryRow(queryGetPrice, item.ProductVariantID).Scan(&price)
		if err != nil {
			return err
		}
		amount := price * item.Qty

		totalAmount += amount

		queryInsertCheckoutItem := `
            INSERT INTO checkout_items (product_variant_id, qty, total, checkout_id, created_at, updated_at)
            VALUES ($1, $2, $3, $4, NOW(), NOW())
        `
		_, err = tx.Exec(queryInsertCheckoutItem, item.ProductVariantID, item.Qty, amount, 0) // `checkout_id` akan diperbarui nanti
		if err != nil {
			return err
		}
	}

	queryInsertCheckout := `
        INSERT INTO checkouts (user_id, total_amount, payment, payment_status, created_at, updated_at)
        VALUES ($1, $2,'COD', 'Paid', NOW(), NOW())
        RETURNING id
    `
	var checkoutID int
	err = tx.QueryRow(queryInsertCheckout, userID, totalAmount).Scan(&checkoutID)
	if err != nil {
		return err
	}

	queryUpdateCheckoutItems := `
        UPDATE checkout_items
        SET checkout_id = $1
        WHERE checkout_id = 0
    `
	_, err = tx.Exec(queryUpdateCheckoutItems, checkoutID)
	if err != nil {
		return err
	}

	queryDeleteCart := `DELETE FROM shopping_carts WHERE user_id = $1`
	_, err = tx.Exec(queryDeleteCart, userID)
	if err != nil {
		return err
	}

	return nil
}
