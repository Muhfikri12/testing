package middleware

import (
	"database/sql"
	"ecommers/database"
	"ecommers/helper"
	"ecommers/util"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type AuthHandler struct {
	Log *zap.Logger
}

func NewAuthHandler(log *zap.Logger) AuthHandler {
	return AuthHandler{Log: log}
}

func (a *AuthHandler) AuthenticateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			a.Log.Error("Missing Token")
			helper.Responses(w, http.StatusUnauthorized, "Missing token", nil)
			return
		}

		config, err := util.ReadConfiguration()
		if err != nil {
			a.Log.Error("failed configurations" + err.Error())
			helper.Responses(w, http.StatusInternalServerError, "failed to read configurations: "+err.Error(), nil)
		}

		db, err := database.InitDB(config)
		if err != nil {
			a.Log.Error("failed database" + err.Error())
			helper.Responses(w, http.StatusInternalServerError, "failed to connect database: "+err.Error(), nil)
		}
		defer db.Close()

		var expiresAt time.Time
		query := `SELECT expired FROM users WHERE token = $1`
		err = db.QueryRow(query, authHeader).Scan(&expiresAt)
		if err != nil {
			if err == sql.ErrNoRows {
				a.Log.Error("Invalid token" + err.Error())
				helper.Responses(w, http.StatusUnauthorized, "Invalid token", nil)
				return
			}
			a.Log.Error("server error" + err.Error())
			http.Error(w, "Server error"+err.Error(), http.StatusInternalServerError)
			return
		}

		if time.Now().After(expiresAt) {
			a.Log.Error("Token expired")
			helper.Responses(w, http.StatusUnauthorized, "Token expired", nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}
