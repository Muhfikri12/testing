package middleware

import (
	"database/sql"
	"ecommers/database"
	"ecommers/helper"
	"ecommers/util"
	"net/http"
	"time"
)

func AuthenticateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helper.Responses(w, http.StatusUnauthorized, "Missing token", nil)
			return
		}

		config, err := util.ReadConfiguration()
		if err != nil {
			helper.Responses(w, http.StatusInternalServerError, "failed to read configurations: "+err.Error(), nil)
		}

		db, err := database.InitDB(config)
		if err != nil {
			helper.Responses(w, http.StatusInternalServerError, "failed to connect database: "+err.Error(), nil)
		}
		defer db.Close()

		var expiresAt time.Time
		query := `SELECT expired FROM users WHERE token = $1`
		err = db.QueryRow(query, authHeader).Scan(&expiresAt)
		if err != nil {
			if err == sql.ErrNoRows {
				helper.Responses(w, http.StatusUnauthorized, "Invalid token", nil)
				return
			}
			http.Error(w, "Server error"+err.Error(), http.StatusInternalServerError)
			return
		}

		if time.Now().After(expiresAt) {
			helper.Responses(w, http.StatusUnauthorized, "Token expired", nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}
