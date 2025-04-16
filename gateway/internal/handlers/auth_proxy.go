package handlers

import (
	"net/http"
	"net/url"

	"github.com/example/get-ur-ghibli/gateway/internal/config"
)

func ProxySignUp(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		proxyRequest(w, r, cfg.AuthServiceURL+"/auth/signup", "POST")
	}
}

func ProxyLogin(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		proxyRequest(w, r, cfg.AuthServiceURL+"/auth/login", "POST")
	}
}

func ProxyCheckFree(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(cfg.AuthServiceURL + "/auth/check-free")
		q := r.URL.Query()
		u.RawQuery = q.Encode()
		proxyRequestCustomURL(w, r, u.String(), "GET")
	}
}

func ProxySetFreeUsed(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(cfg.AuthServiceURL + "/auth/set-free-used")
		q := r.URL.Query()
		u.RawQuery = q.Encode()
		proxyRequestCustomURL(w, r, u.String(), "POST")
	}
}
