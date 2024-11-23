package auth

import (
	"database/sql"
	"ecommers/model"
	"errors"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type AuthRepository struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewAuthRepository(db *sql.DB, Log *zap.Logger) AuthRepository {
	return AuthRepository{
		DB:     db,
		Logger: Log,
	}
}

func (a *AuthRepository) Login(login *model.Login) error {
	query := `SELECT id FROM users WHERE (email=$1 OR phone=$1) AND password=$2`
	err := a.DB.QueryRow(query, login.Email, login.Password).Scan(&login.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid email or password")
		}
		return err
	}

	token := uuid.New().String()

	expiration := time.Now().UTC().Add(1 * time.Hour)
	updateQuery := `UPDATE users SET token=$1, expired=$2 WHERE id=$3`
	_, err = a.DB.Exec(updateQuery, token, expiration, login.ID)
	if err != nil {
		return err
	}

	login.Token = token
	login.Email = ""
	login.Password = ""

	return nil
}
