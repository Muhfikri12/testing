package users

import (
	"ecommers/model"
)

func (u *UsersRepository) GetListAddress(token string) (*[]model.Users, error) {
	query := `SELECT a.address FROM users u JOIN addresses a ON a.user_id = u.id WHERE u.token = $1`

	rows, err := u.DB.Query(query, token)
	if err != nil {
		u.Logger.Error("Error from query GetListAddress: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	users := []model.Users{}

	for rows.Next() {
		user := model.Users{}
		if err := rows.Scan(&user.Address); err != nil {
			u.Logger.Error("Error from Scan GetListAddress: " + err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}

func (u *UsersRepository) GetDetailUser(token string, user *model.Users) error {

	query := `SELECT u.id, u.name, u.username, u.email, a.address, u.phone
		JOIN address a ON u.id = a.user_id
		WHERE u.token = $1`

	if err := u.DB.QueryRow(query, token).Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Address, &user.Phone); err != nil {
		u.Logger.Error("Error from query GetDetailUser: " + err.Error())
		return err
	}

	return nil
}
