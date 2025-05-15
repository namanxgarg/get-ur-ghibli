package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"

	"github.com/example/get-ur-ghibli/upload-service/internal/handlers"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zlog.With().Str("service", "upload-service").Logger()

	r := mux.NewRouter()
	r.HandleFunc("/upload", handlers.HandleUpload).Methods("POST")

	logger.Info().Msg("Upload service listening on :8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
