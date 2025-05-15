package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/example/get-ur-ghibli/auth-service/internal/config"
	"github.com/example/get-ur-ghibli/auth-service/internal/models"
	"github.com/example/get-ur-ghibli/auth-service/internal/repository"
	"github.com/example/get-ur-ghibli/auth-service/internal/token"
	"github.com/jinzhu/gorm"
)

func SignUpHandler(db *gorm.DB, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			Role     string `json:"role"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userRepo := repository.NewUserRepository(db)
		existing, _ := userRepo.FindByEmail(req.Email)
		if existing != nil {
			http.Error(w, "Email already taken", http.StatusConflict)
			return
		}

		// Hash password
		hash := sha256.Sum256([]byte(req.Password))
		user := &models.User{
			Email:        req.Email,
			PasswordHash: hex.EncodeToString(hash[:]),
			Role:         req.Role,
		}

		if err := userRepo.CreateUser(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User created successfully"))
	}
}

func LoginHandler(db *gorm.DB, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userRepo := repository.NewUserRepository(db)
		user, err := userRepo.FindByEmail(req.Email)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Check password
		hash := sha256.Sum256([]byte(req.Password))
		if hex.EncodeToString(hash[:]) != user.PasswordHash {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Generate JWT
		tokenString, err := token.GenerateTokenWithRole(user.Email, user.Role, cfg.JWTSecret)
		if err != nil {
			http.Error(w, "Failed to create token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}

// Check if user has used free image
func CheckFreeImageHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// We'll parse the email from a header or something in real usage
		// For simplicity, let's read from query param
		email := r.URL.Query().Get("email")
		if email == "" {
			http.Error(w, "Email required", http.StatusBadRequest)
			return
		}

		userRepo := repository.NewUserRepository(db)
		user, err := userRepo.FindByEmail(email)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		used := user.HasUsedFree
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]bool{"hasUsedFree": used})
	}
}

// Mark free image as used
func SetFreeUsedHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		if email == "" {
			http.Error(w, "Email required", http.StatusBadRequest)
			return
		}

		userRepo := repository.NewUserRepository(db)
		user, err := userRepo.FindByEmail(email)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		user.HasUsedFree = true
		if err := userRepo.UpdateUser(user); err != nil {
			http.Error(w, "Failed to update user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Free image marked as used"))
	}
}
