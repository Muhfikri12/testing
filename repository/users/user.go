package users

import "ecommers/model"

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
