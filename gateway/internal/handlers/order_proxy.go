package handlers

import (
	"net/http"

	"github.com/example/get-ur-ghibli/gateway/internal/config"
	"github.com/gorilla/mux"
)

func ProxyCreateOrder(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		target := cfg.OrderServiceURL + "/orders"
		proxyRequest(w, r, target, "POST")
	}
}

func ProxyPayOrder(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		orderID := vars["orderID"]
		target := cfg.OrderServiceURL + "/orders/" + orderID + "/pay"
		proxyRequest(w, r, target, "POST")
	}
}

func ProxyGetOrder(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		orderID := vars["orderID"]
		target := cfg.OrderServiceURL + "/orders/" + orderID
		proxyRequest(w, r, target, "GET")
	}
}
