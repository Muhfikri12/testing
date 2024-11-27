package auth_test

import (
	"ecommers/model"
	"ecommers/repository/auth"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestAuthRepository_Register(t *testing.T) {
	// Buat mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Inisialisasi logger dan repository
	logger := zap.NewNop() // Menggunakan logger kosong untuk test
	authRepo := auth.NewAuthRepository(db, logger)

	t.Run("successfully registers a user", func(t *testing.T) {

		user := &model.Register{
			Name:     "John Doe",
			Phone:    "1234567890",
			Email:    "john@example.com",
			Password: "securepassword",
		}

		mock.ExpectQuery(`INSERT INTO users`).
			WithArgs(
				user.Name,
				sqlmock.AnyArg(),
				user.Phone,
				user.Email,
				user.Password,
			).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		// Eksekusi fungsi Register
		err := authRepo.Register(user)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, 1, int(user.ID))
		assert.NotEmpty(t, user.Username)
	})

	t.Run("fails due to duplicate email or phone", func(t *testing.T) {
		// Data input user
		user := &model.Register{
			Name:     "Jane Doe",
			Phone:    "0987654321",
			Email:    "jane@example.com",
			Password: "securepassword",
		}

		mock.ExpectQuery(`INSERT INTO users`).
			WithArgs(
				user.Name,
				sqlmock.AnyArg(),
				user.Phone,
				user.Email,
				user.Password,
			).
			WillReturnError(errors.New("duplicate key value violates unique constraint"))

		err := authRepo.Register(user)

		assert.Error(t, err)
		assert.Equal(t, "email or phone already in use", err.Error())
	})

}
