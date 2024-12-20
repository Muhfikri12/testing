package products

import (
	"fmt"
)

func (w *ProductRepository) CreateWishlisht(token string, id int) error {
	querySelectProduct := `SELECT id FROM product_varians WHERE id=$1`
	err := w.DB.QueryRow(querySelectProduct, id).Scan(&id)
	if err != nil {
		w.Logger.Error("Product not found: " + err.Error())
		return fmt.Errorf("product not found")
	}

	queryDetectProduct := `
        SELECT COUNT(*) 
        FROM wishlists 
        WHERE product_variant_id=$1 
          AND user_id = (SELECT id FROM users WHERE token=$2)
    `
	var count int
	if err := w.DB.QueryRow(queryDetectProduct, id, token).Scan(&count); err != nil {
		w.Logger.Error("Failed to Query Data " + err.Error())
		return fmt.Errorf("failed to query data: %w", err)
	}

	if count > 0 {
		w.Logger.Debug("Product already exists in wishlist")
		return fmt.Errorf("product already exists in wishlist")
	}

	query := `INSERT INTO wishlists(product_variant_id, user_id, created_at) VALUES($1,(SELECT id FROM users WHERE token=$2), NOW()) RETURNING id`
	if err := w.DB.QueryRow(query, id, token).Scan(&id); err != nil {
		w.Logger.Error("Error from repository: " + err.Error())
		return err
	}
	return nil
}

func (w *ProductRepository) DeleteWishlist(id int, token string) error {

	query := `DELETE FROM wishlists WHERE product_variant_id=$1 AND user_id = (SELECT id FROM users WHERE token =$2)`

	result, err := w.DB.Exec(query, id, token)
	if err != nil {
		w.Logger.Error("Error executing Delete: " + err.Error())
		return fmt.Errorf("failed to delete wishlist")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		w.Logger.Error("Error fetching rows affected: " + err.Error())
		return fmt.Errorf("failed to fetch rows affected")
	}

	if rowsAffected == 0 {
		w.Logger.Debug("No wishlist entry found to delete")
		return fmt.Errorf("no matching wishlist entry found")
	}

	return nil
}
