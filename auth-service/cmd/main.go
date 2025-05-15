package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	zlog "github.com/rs/zerolog/log"

	"github.com/example/get-ur-ghibli/auth-service/internal/config"
	"github.com/example/get-ur-ghibli/auth-service/internal/database"
	"github.com/example/get-ur-ghibli/auth-service/internal/handlers"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zlog.With().Str("service", "auth-service").Logger()

	cfg := config.LoadConfigFromEnv()

	db, err := database.InitDB(cfg)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer db.Close()

	// Migrate
	database.AutoMigrate(db)

	r := mux.NewRouter()
	r.HandleFunc("/auth/signup", handlers.SignUpHandler(db, cfg)).Methods("POST")
	r.HandleFunc("/auth/login", handlers.LoginHandler(db, cfg)).Methods("POST")
	r.HandleFunc("/auth/check-free", handlers.CheckFreeImageHandler(db)).Methods("GET")
	r.HandleFunc("/auth/set-free-used", handlers.SetFreeUsedHandler(db)).Methods("POST")

	logger.Info().Msgf("Auth service listening on :8081")
	http.ListenAndServe(":8081", r)
}
