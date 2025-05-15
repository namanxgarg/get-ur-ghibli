package router

import (
    "github.com/gorilla/mux"
    "github.com/example/get-ur-ghibli/gateway/internal/config"
    "github.com/example/get-ur-ghibli/gateway/internal/handlers"
    "github.com/example/get-ur-ghibli/gateway/internal/middlewares"
)

func InitRoutes(r *mux.Router, cfg *config.Config) {
    // Auth routes (public)
    r.HandleFunc("/api/auth/signup", handlers.ProxySignUp(cfg)).Methods("POST")
    r.HandleFunc("/api/auth/login", handlers.ProxyLogin(cfg)).Methods("POST")

    // Check free usage
    r.HandleFunc("/api/auth/check-free", handlers.ProxyCheckFree(cfg)).Methods("GET")
    r.HandleFunc("/api/auth/set-free-used", handlers.ProxySetFreeUsed(cfg)).Methods("POST")

    // Upload route (protected, or maybe public, your choice)
    r.HandleFunc("/api/upload", middlewares.JWTAuth(handlers.ProxyUpload(cfg), cfg)).Methods("POST")

    // Ghibli generation
    r.HandleFunc("/api/ghibli/free/{imageID}", middlewares.JWTAuth(handlers.ProxyGenerateFree(cfg), cfg)).Methods("GET")
    r.HandleFunc("/api/ghibli/paid/{imageID}", middlewares.JWTAuth(handlers.ProxyGeneratePaid(cfg), cfg)).Methods("GET")

    // Orders
    r.HandleFunc("/api/orders", middlewares.JWTAuth(handlers.ProxyCreateOrder(cfg), cfg)).Methods("POST")
    r.HandleFunc("/api/orders/{orderID}/pay", middlewares.JWTAuth(handlers.ProxyPayOrder(cfg), cfg)).Methods("POST")
    r.HandleFunc("/api/orders/{orderID}", middlewares.JWTAuth(handlers.ProxyGetOrder(cfg), cfg)).Methods("GET")
}
