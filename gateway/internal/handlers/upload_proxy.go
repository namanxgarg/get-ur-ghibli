package handlers

import (
    "github.com/example/get-ur-ghibli/gateway/internal/config"
    "net/http"
)

func ProxyUpload(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        target := cfg.UploadServiceURL + "/upload"
        proxyRequest(w, r, target, "POST")
    }
}
