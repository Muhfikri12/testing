package auth

import (
	"database/sql"
	"ecommers/model"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/exp/rand"
)

type AuthRepositoryInterface interface {
	Login(login *model.Login) error
	Register(user *model.Register) error
}

type AuthRepository struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewAuthRepository(db *sql.DB, Log *zap.Logger) AuthRepositoryInterface {
	return &AuthRepository{
		DB:     db,
		Logger: Log,
	}
}

func (a *AuthRepository) Login(login *model.Login) error {

	query := `SELECT id FROM users WHERE email=$1 OR phone=$1`
	var userID int
	err := a.DB.QueryRow(query, login.Email).Scan(&userID)

	if err != nil {
		if err == sql.ErrNoRows {
			a.Logger.Debug("Invalid Phone Or Email: " + err.Error())
			return errors.New("email or phone not found")
		}
		return err
	}

	passwordQuery := `SELECT id FROM users WHERE (email=$1 OR phone=$1) AND password=$2`
	err = a.DB.QueryRow(passwordQuery, login.Email, login.Password).Scan(&login.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			a.Logger.Debug("Invalid Password: " + err.Error())
			return errors.New("password salah")
		}
		return err
	}

	token := uuid.New().String()
	expiration := time.Now().UTC().Add(1 * time.Hour)

	updateQuery := `UPDATE users SET token=$1, expired=$2 WHERE id=$3`
	_, err = a.DB.Exec(updateQuery, token, expiration, login.ID)
	if err != nil {
		a.Logger.Error("Error from Login repository: " + err.Error())
		return err
	}

	login.Token = token
	login.Email = ""
	login.Password = ""

	return nil
}

func (a *AuthRepository) Register(user *model.Register) error {

	// Generate username berdasarkan nama pengguna
	user.Username = strings.ReplaceAll(strings.ToLower(user.Name), " ", "")
	user.Username += fmt.Sprintf("%04d", rand.Intn(10000)) // Tambahkan 4 digit angka acak

	// Query untuk menambahkan user baru ke database
	query := `
		INSERT INTO users (name, username, phone, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id
	`

	// Eksekusi query
	err := a.DB.QueryRow(query, user.Name, user.Username, user.Phone, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		// Periksa apakah error disebabkan oleh duplikasi data (contoh: email/username sudah digunakan)
		if strings.Contains(err.Error(), "duplicate") {
			a.Logger.Warn("Duplicate entry detected during registration: " + err.Error())
			return errors.New("email or phone already in use")
		}

		// Log error lain dan kembalikan error tersebut
		a.Logger.Error("Error from register repository: " + err.Error())
		return err
	}

	// Berhasil mendaftarkan pengguna
	a.Logger.Info(fmt.Sprintf("User registered successfully with ID: %d", user.ID))
	return nil
}
