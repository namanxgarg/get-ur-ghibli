package middlewares

import (
    "net/http"
    "strings"

    "github.com/example/get-ur-ghibli/gateway/internal/config"
    "github.com/golang-jwt/jwt/v4"
)

func JWTAuth(next http.HandlerFunc, cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing auth token", http.StatusUnauthorized)
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, "Invalid auth header", http.StatusUnauthorized)
            return
        }
        tokenStr := parts[1]

        token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
            return []byte(cfg.JWTSecret), nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    }
}
