package users

import (
	"ecommers/model"
	"fmt"
)

func (u *UsersRepository) AddAddress(token string, address *model.Addresses) error {

	var addressCount int
	queryCheckAddress := `SELECT COUNT(*) FROM addresses WHERE user_id = (SELECT id FROM users WHERE token=$1)`
	if err := u.DB.QueryRow(queryCheckAddress, token).Scan(&addressCount); err != nil {
		u.Logger.Error("Failed to check existing addresses: " + err.Error())
		return err
	}

	isMain := addressCount == 0

	queryInsert := `INSERT INTO addresses (address, user_id, is_main, created_at, updated_at) 
                    VALUES ($1, (SELECT id FROM users WHERE token=$2), $3, NOW(), NOW()) RETURNING id`
	if err := u.DB.QueryRow(queryInsert, address.Address, token, isMain).Scan(&address.ID); err != nil {
		u.Logger.Error("Failed to insert address: " + err.Error())
		return err
	}

	return nil
}

func (u *UsersRepository) UpdateAddress(token string, id int, address *model.Addresses) error {

	query := `UPDATE addresses 
              SET address=$1, updated_at=NOW() 
              WHERE id = $2 AND user_id = (SELECT id FROM users WH	ERE token = $3)`

	_, err := u.DB.Exec(query, address.Address, id, token)
	if err != nil {
		u.Logger.Error("Error updating address: " + err.Error())
		return err
	}

	return nil
}

func (u *UsersRepository) DeleteAddress(token string, id int) error {

	query := `DELETE FROM addresses 
	WHERE user_id = (SELECT id FROM users WHERE token = $1) AND id=$2`

	result, err := u.DB.Exec(query, token, id)
	if err != nil {
		u.Logger.Error("Error updating address: " + err.Error())
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		u.Logger.Error("Error fetching rows affected: " + err.Error())
		return fmt.Errorf("failed to fetch rows affected")
	}

	if rowsAffected == 0 {
		u.Logger.Debug("No Address entry found to delete")
		return fmt.Errorf("no matching Address entry found")
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

func (u *UsersRepository) GetListAddress(token string) (*[]model.Addresses, error) {

	var isMain bool
	query := `SELECT address, is_main FROM addresses WHERE user_id = (SELECT id FROM users WHERE token=$1) `

	rows, err := u.DB.Query(query, token)
	if err != nil {
		u.Logger.Error("Error from query GetListAddress: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	users := []model.Addresses{}

	for rows.Next() {
		user := model.Addresses{}
		if err := rows.Scan(&user.Address, &isMain); err != nil {
			u.Logger.Error("Error from Scan GetListAddress: " + err.Error())
			return nil, err
		}

		if isMain {
			user.Status = "default"
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		u.Logger.Warn("No addresses found for the given token")
		return nil, fmt.Errorf("no addresses found for token: %s", token)
	}

	return &users, nil
}
