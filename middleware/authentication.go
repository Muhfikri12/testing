package middleware

import (
	"database/sql"
	"ecommers/database"
	"ecommers/helper"
	"ecommers/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct {
	Log *zap.Logger
}

func NewAuthHandler(log *zap.Logger) AuthHandler {
	return AuthHandler{Log: log}
}

func (a *AuthHandler) AuthenticateGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			a.Log.Error("Missing Token")
			helper.ResponsesJson(c, http.StatusUnauthorized, "Missing token", nil)
			c.Abort()
			return
		}

		// Baca konfigurasi
		config, err := util.ReadConfiguration()
		if err != nil {
			a.Log.Error("Failed to read configurations: " + err.Error())
			helper.ResponsesJson(c, http.StatusInternalServerError, "Failed to read configurations", nil)
			c.Abort()
			return
		}

		// Inisialisasi database
		db, err := database.InitDB(config)
		if err != nil {
			a.Log.Error("Failed to connect to database: " + err.Error())
			helper.ResponsesJson(c, http.StatusInternalServerError, "Failed to connect to database", nil)
			c.Abort()
			return
		}
		defer db.Close()

		// Validasi token di database
		var expiresAt time.Time
		query := `SELECT expired FROM users WHERE token = $1`
		err = db.QueryRow(query, authHeader).Scan(&expiresAt)
		if err != nil {
			if err == sql.ErrNoRows {
				a.Log.Error("Invalid token")
				helper.ResponsesJson(c, http.StatusUnauthorized, "Invalid token", nil)
				c.Abort()
				return
			}
			a.Log.Error("Server error: " + err.Error())
			helper.ResponsesJson(c, http.StatusInternalServerError, "Server error", nil)
			c.Abort()
			return
		}

		// Periksa apakah token sudah expired
		if time.Now().UTC().After(expiresAt) {
			a.Log.Error("Token expired")
			helper.ResponsesJson(c, http.StatusUnauthorized, "Token expired", nil)
			c.Abort()
			return
		}

		// Token valid, lanjutkan ke handler berikutnya
		c.Next()
	}
}
