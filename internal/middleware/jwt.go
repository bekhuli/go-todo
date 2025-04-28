package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/bekhuli/go-todo/config"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserKey contextKey = "userID"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := config.Envs

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "No token provided", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		UserID, ok := claims["user_id"].(float64)
		if !ok {
			http.Error(w, "Invalid user_id in token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, int(UserID))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
