package handlers

import (
    "github.com/example/get-ur-ghibli/gateway/internal/config"
    "github.com/gorilla/mux"
    "net/http"
)

func ProxyGenerateFree(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        imageID := vars["imageID"]
        target := cfg.GhibliServiceURL + "/generate/free/" + imageID
        proxyRequest(w, r, target, "GET")
    }
}

func ProxyGeneratePaid(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        imageID := vars["imageID"]
        target := cfg.GhibliServiceURL + "/generate/paid/" + imageID
        proxyRequest(w, r, target, "GET")
    }
}
