package checkout

import (
	"ecommers/model"
	"fmt"
)

func (c *CheckoutsRepository) ProcessCheckout(token string) (*model.Checkouts, error) {

	tx, err := c.DB.Begin()
	if err != nil {
		return nil, err
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
        WHERE user_id = (SELECT id FROM users WHERE token=$1) AND deleted_at IS NULL
    `
	rows, err := tx.Query(queryCartItems, token)
	if err != nil {
		return nil, err
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
			return nil, err
		}
		cartItems = append(cartItems, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(cartItems) == 0 {
		return nil, fmt.Errorf("shopping cart is empty")
	}

	var totalAmount int
	for _, item := range cartItems {
		var price, stock int

		queryGetPriceAndStock := `
        SELECT p.price - (p.price * p.discount / 100) as price, pv.stock
        FROM product_varians pv
        JOIN products p ON pv.product_id = p.id
        WHERE pv.id = $1
    `
		err = tx.QueryRow(queryGetPriceAndStock, item.ProductVariantID).Scan(&price, &stock)
		if err != nil {
			return nil, err
		}

		if stock < item.Qty {
			return nil, fmt.Errorf("insufficient stock for product_variant_id: %d", item.ProductVariantID)
		}

		queryUpdateStock := `
        UPDATE product_varians
        SET stock = stock - $1, updated_at = NOW()
        WHERE id = $2
    `
		_, err = tx.Exec(queryUpdateStock, item.Qty, item.ProductVariantID)
		if err != nil {
			return nil, err
		}

		amount := price * item.Qty
		totalAmount += amount

		queryInsertCheckoutItem := `
        INSERT INTO checkout_items (product_variant_id, qty, total, checkout_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, NOW(), NOW())
    `
		_, err = tx.Exec(queryInsertCheckoutItem, item.ProductVariantID, item.Qty, amount, 0)
		if err != nil {
			return nil, err
		}
	}

	queryInsertCheckout := `
        INSERT INTO checkouts (user_id, total_amount, payment, payment_status, created_at, updated_at)
        VALUES ((SELECT id FROM users WHERE token=$1), $2,'COD', 'Paid', NOW(), NOW())
        RETURNING id
    `
	var checkoutID int
	err = tx.QueryRow(queryInsertCheckout, token, totalAmount).Scan(&checkoutID)
	if err != nil {
		return nil, err
	}

	queryUpdateCheckoutItems := `
        UPDATE checkout_items
        SET checkout_id = $1
        WHERE checkout_id = 0
    `
	_, err = tx.Exec(queryUpdateCheckoutItems, checkoutID)
	if err != nil {
		return nil, err
	}

	queryDeleteCart := `DELETE FROM shopping_carts WHERE user_id = (SELECT id FROM users WHERE token=$1)`
	_, err = tx.Exec(queryDeleteCart, token)
	if err != nil {
		return nil, err
	}

	checkout := model.Checkouts{
		Users:    &model.Users{},
		Products: &[]model.Products{},
	}

	querySelectDetailOrder := `
		SELECT u.name, c.total_amount, c.payment, c.payment_status 
			FROM checkouts c
			JOIN users u ON c.user_id = u.id
			WHERE c.id=$1`

	if err := tx.QueryRow(querySelectDetailOrder, checkoutID).Scan(&checkout.Users.Name, &checkout.TotalAmount, &checkout.Payment, &checkout.PaymentStatus); err != nil {
		c.Logger.Error("Failed to Get Detail Order: " + err.Error())
		return nil, err
	}

	products := []model.Products{}
	queryGetDetailProduct := `
		SELECT p.name, c.qty, c.total, pv.size, pv.color
			FROM checkout_items c
			JOIN product_varians pv ON c.product_variant_id = pv.id
			JOIN products p ON pv.product_id = p.id
			WHERE c.checkout_id = $1`

	rowsProduct, err := tx.Query(queryGetDetailProduct, checkoutID)
	if err != nil {
		c.Logger.Error("Failed to Query Detail Product: " + err.Error())
		return nil, err
	}
	defer rowsProduct.Close()

	for rowsProduct.Next() {
		product := model.Products{}
		if err := rowsProduct.Scan(&product.Name, &product.Qty, &product.Amount, &product.Size, &product.Color); err != nil {
			c.Logger.Error("Failed to Scan Detail Product: " + err.Error())
		}
		products = append(products, product)
	}

	checkout.Products = &products

	return &checkout, nil
}
