package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	// zlog "github.com/rs/zerolog"

	zlog "github.com/rs/zerolog/log"

	"github.com/example/get-ur-ghibli/ghibli-service/internal/handlers"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zlog.With().Str("service", "ghibli-service").Logger()

	handlers.StartKafkaConsumer()
	r := mux.NewRouter()
	// Endpoint for 1 free image
	r.HandleFunc("/generate/free/{imageID}", handlers.GenerateFreeHandler).Methods("GET")

	// Endpoint for 10 paid images
	r.HandleFunc("/generate/paid/{imageID}", handlers.GeneratePaidHandler).Methods("GET")

	r.HandleFunc("/job-status", handlers.JobStatusHandler).Methods("GET")

	logger.Info().Msg("Ghibli service listening on :8083")
	log.Fatal(http.ListenAndServe(":8083", r))
}
