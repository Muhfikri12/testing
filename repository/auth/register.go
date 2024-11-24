package auth

import (
	"ecommers/model"
	"fmt"
	"strings"

	"golang.org/x/exp/rand"
)

func (u *AuthRepository) Register(user *model.Register) error {

	user.Username = strings.ReplaceAll(user.Name, " ", "")
	user.Username += fmt.Sprintf("%d", rand.Intn(1000))

	checkQuery := `SELECT COUNT(*) FROM users WHERE email = $1 OR phone = $2`
	var count int
	err := u.DB.QueryRow(checkQuery, user.Email, user.Phone).Scan(&count)
	if err != nil {
		u.Logger.Error("Error checking email/phone existence: " + err.Error())
		return err
	}

	if count > 0 {
		return fmt.Errorf("email or phone already in use")
	}

	query := `INSERT INTO users(name, username, phone, email, password, created_at, updated_at) values($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING id`

	if err := u.DB.QueryRow(query, user.Name, user.Username, user.Phone, user.Email, user.Password).Scan(&user.ID); err != nil {
		u.Logger.Error("Error from register repository: " + err.Error())
		return err
	}

	return nil
}
