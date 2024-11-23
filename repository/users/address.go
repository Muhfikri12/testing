package users

import "ecommers/model"

func (u *UsersRepository) AddAddress(token string, address *model.Addresses) error {

	var userID int

	queryGetUserID := `SELECT id FROM users WHERE token = $1`
	if err := u.DB.QueryRow(queryGetUserID, token).Scan(&userID); err != nil {
		u.Logger.Error("Failed to insert address: " + err.Error())
		return err
	}

	query := `INSERT INTO addresses (address, user_id, created_at, updated_at) VALUES ($1, $2, NOW(), NOW()) RETURNING id`

	if err := u.DB.QueryRow(query, address.Address, userID).Scan(&address.ID); err != nil {
		u.Logger.Error("Failed to insert address: " + err.Error())
		return err
	}

	return nil
}

func (u *UsersRepository) UpdateAddress(token string, address *model.Addresses) error {

	query := `UPDATE addresses 
              SET address=$1, updated_at=NOW() 
              WHERE id = $2 AND user_id = (SELECT id FROM users WHERE token = $3)`

	_, err := u.DB.Exec(query, address.Address, address.ID, token)
	if err != nil {
		u.Logger.Error("Error updating address: " + err.Error())
		return err
	}

	return nil
}

func (u *UsersRepository) DeleteAddress(token string, id int) error {

	query := `DELETE FROM addresses 
	WHERE user_id = (SELECT id FROM users WHERE token = $1) AND id=$2`

	_, err := u.DB.Exec(query, token, id)
	if err != nil {
		u.Logger.Error("Error updating address: " + err.Error())
		return err
	}

	return nil
}

func (u *UsersRepository) SetDefault(token string, id int) error {

	tx, err := u.DB.Begin()
	if err != nil {
		u.Logger.Error("Error starting transaction: " + err.Error())
		return err
	}

	queryUnset := `UPDATE addresses 
                   SET is_main = false, updated_at = NOW() 
                   WHERE user_id = (SELECT id FROM users WHERE token = $1)`
	_, err = tx.Exec(queryUnset, token)
	if err != nil {
		u.Logger.Error("Error unsetting default addresses: " + err.Error())
		tx.Rollback()
		return err
	}

	querySet := `UPDATE addresses 
                 SET is_main = true, updated_at = NOW() 
                 WHERE id = $1 AND user_id = (SELECT id FROM users WHERE token = $2)`
	_, err = tx.Exec(querySet, id, token)
	if err != nil {
		u.Logger.Error("Error setting default address: " + err.Error())
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		u.Logger.Error("Error committing transaction: " + err.Error())
		return err
	}

	return nil
}
