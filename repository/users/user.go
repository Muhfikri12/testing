package users

import (
	"database/sql"
	"ecommers/model"
	"fmt"

	"go.uber.org/zap"
)

type UsersRepoInterface interface {
	GetDetailUser(token string) (*model.Users, error)
	UpdateUserData(token string, user *model.Users) error
	AddAddress(token string, address *model.Addresses) error
	UpdateAddress(token string, id int, address *model.Addresses) error
	DeleteAddress(token string, id int) error
	SetDefault(token string, id int) error
	GetListAddress(token string) (*[]model.Addresses, error)
}

type UsersRepository struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewUsersRepository(db *sql.DB, Log *zap.Logger) UsersRepoInterface {
	return &UsersRepository{
		DB:     db,
		Logger: Log,
	}
}

func (u *UsersRepository) GetDetailUser(token string) (*model.Users, error) {

	user := model.Users{
		Address: &model.Addresses{},
	}
	var isMain bool
	query := `SELECT u.id, u.name, u.email, a.address, a.is_main, u.phone
		FROM users u
		JOIN addresses a ON a.user_id = u.id
		WHERE u.token = $1 AND is_main = true`

	if err := u.DB.QueryRow(query, token).Scan(&user.ID, &user.Name, &user.Email, &user.Address.Address, &isMain, &user.Phone); err != nil {
		u.Logger.Error("Error from query GetDetailUser: " + err.Error())
		return nil, err
	}

	if isMain {
		user.Address.Status = "default"
	}

	return &user, nil
}

func (u *UsersRepository) UpdateUserData(token string, user *model.Users) error {
	if user == nil {
		return fmt.Errorf("user cannot be nil")
	}

	trx, err := u.DB.Begin()
	if err != nil {
		u.Logger.Error("Error starting transaction: " + err.Error())
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			if trx != nil {
				trx.Rollback()
			}
			panic(p)
		} else if err != nil && trx != nil {
			trx.Rollback()
		} else if trx != nil {
			trx.Commit()
		}
	}()

	query := `UPDATE users 
        SET name = $1, phone = $2, email = $3, updated_at = NOW()
        WHERE token = $4`
	_, err = trx.Exec(query, user.Name, user.Phone, user.Email, token)
	if err != nil {
		u.Logger.Error("Error updating user: " + err.Error())
		return err
	}

	if user.CurrentPassword != nil && user.NewPassword != "" {

		var currentPassword string
		query = `SELECT password FROM users WHERE token = $1`
		err = trx.QueryRow(query, token).Scan(&currentPassword)
		if err != nil {
			u.Logger.Error("Failed to fetch current password: " + err.Error())
			return err
		}

		if *user.CurrentPassword != currentPassword {
			u.Logger.Error("Invalid current password")
			return fmt.Errorf("invalid current password")
		}

		if len(user.NewPassword) < 8 {
			u.Logger.Error("New password is too short")
			return fmt.Errorf("new password must be at least 8 characters")
		}

		query = `UPDATE users SET password = $1, updated_at = NOW() WHERE token = $2`
		_, err = trx.Exec(query, user.NewPassword, token)
		if err != nil {
			u.Logger.Error("Failed to update password: " + err.Error())
			return err
		}
	} else if user.NewPassword != "" {

		u.Logger.Error("Current password is required to update the password")
		return fmt.Errorf("current password is required to update the password")
	}

	return nil
}
