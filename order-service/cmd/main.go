package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	zlog "github.com/rs/zerolog/log"

	"github.com/example/get-ur-ghibli/order-service/internal/config"
	"github.com/example/get-ur-ghibli/order-service/internal/database"
	"github.com/example/get-ur-ghibli/order-service/internal/handlers"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zlog.With().Str("service", "order-service").Logger()

	cfg := config.LoadConfigFromEnv()

	db, err := database.InitDB(cfg)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer db.Close()

	// Auto migrate
	database.AutoMigrate(db)

	r := mux.NewRouter()
	// Create an order
	r.HandleFunc("/orders", handlers.CreateOrderHandler(db)).Methods("POST")
	// Pay for an order
	r.HandleFunc("/orders/{orderID}/pay", handlers.PayOrderHandler(db)).Methods("POST")
	// Track an order
	r.HandleFunc("/orders/{orderID}", handlers.GetOrderHandler(db)).Methods("GET")

	logger.Info().Msg("Order service listening on :8084")
	log.Fatal(http.ListenAndServe(":8084", r))
}
