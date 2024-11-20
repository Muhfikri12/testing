package products

import (
	"ecommers/model"
	"fmt"
	"time"
)

func (w *ProductRepository) CreateWishlisht(wishlists *model.Wishlists) error {
	today := time.Now()

	wishlists.Timestamps.Created_at = &today

	query := `INSERT INTO wishlists(product_id, user_id, created_at) VALUES($1, $2, $3) RETURNING id`
	if err := w.DB.QueryRow(query, wishlists.ProductID, wishlists.UserID, wishlists.Timestamps.Created_at).Scan(&wishlists.ID); err != nil {
		w.Logger.Error("Error from repository: " + err.Error())
		return err
	}

	return nil
}

func (w *ProductRepository) DeleteWishlist(id int) error {

	query := `UPDATE wishlists SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL`

	now := time.Now()

	result, err := w.DB.Exec(query, now, id)
	if err != nil {
		w.Logger.Error("Error executing soft delete: " + err.Error())
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		w.Logger.Error("Error checking rows affected: " + err.Error())
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no record found to delete or already deleted")
	}

	return nil
}
