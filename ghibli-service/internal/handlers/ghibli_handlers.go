package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/example/get-ur-ghibli/ghibli-service/internal/generation"
	"github.com/gorilla/mux"
)

// For demonstration, we do not check user or payments here
// but you *could* confirm user info via query param or header

func GenerateFreeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageID := vars["imageID"]

	// Generate a single image
	result := generation.GenerateMock(imageID, 1)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func GeneratePaidHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageID := vars["imageID"]

	// Generate 10 images
	result := generation.GenerateMock(imageID, 10)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
