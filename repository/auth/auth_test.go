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
		// Data input user valid
		user := &model.Register{
			Name:     "John Doe",
			Phone:    "1234567890",
			Email:    "john@example.com",
			Password: "securepassword",
		}

		mock.ExpectQuery(`INSERT INTO users`).
			WithArgs(
				user.Name,
				sqlmock.AnyArg(), // Username yang di-generate akan dinamis
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

		// Mock query expectation with duplicate error
		mock.ExpectQuery(`INSERT INTO users`).
			WithArgs(
				user.Name,
				sqlmock.AnyArg(),
				user.Phone,
				user.Email,
				user.Password,
			).
			WillReturnError(errors.New("duplicate key value violates unique constraint"))

		// Eksekusi fungsi Register
		err := authRepo.Register(user)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, "email or phone already in use", err.Error())
	})

	t.Run("fails due to generic database error", func(t *testing.T) {
		// Data input user
		user := &model.Register{
			Name:     "Alice Doe",
			Phone:    "5678901234",
			Email:    "alice@example.com",
			Password: "securepassword",
		}

		// Mock query expectation with generic error
		mock.ExpectQuery(`INSERT INTO users`).
			WithArgs(
				user.Name,
				sqlmock.AnyArg(),
				user.Phone,
				user.Email,
				user.Password,
			).
			WillReturnError(errors.New("generic database error"))

		// Eksekusi fungsi Register
		err := authRepo.Register(user)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, "generic database error", err.Error())
	})
}
