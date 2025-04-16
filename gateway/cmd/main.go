package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	zlog "github.com/rs/zerolog/log"

	"github.com/example/get-ur-ghibli/gateway/internal/config"
	"github.com/example/get-ur-ghibli/gateway/internal/router"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zlog.With().Str("service", "gateway").Logger()

	cfg := config.LoadConfigFromEnv()
	r := mux.NewRouter()

	router.InitRoutes(r, cfg)

	logger.Info().Msg("Gateway listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
