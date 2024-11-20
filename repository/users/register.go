package users

import (
	"ecommers/model"
	"fmt"
	"strings"

	"golang.org/x/exp/rand"
)

func (u *UsersRepository) Register(user *model.Users) error {

	user.Username = strings.ReplaceAll(user.Name, " ", "")

	user.Username += fmt.Sprintf("%d", rand.Intn(1000))

	query := `INSERT INTO users(name, username, address, phone, email, password) values($1, $2, $3, $4, $5, $6) RETURNING id`

	if err := u.DB.QueryRow(query, user.Name, user.Username, user.Address, user.Phone, user.Email, user.Password).Scan(&user.ID); err != nil {
		u.Logger.Error("Error from register repository: " + err.Error())
		return err
	}

	user.Password = ""

	return nil
}
