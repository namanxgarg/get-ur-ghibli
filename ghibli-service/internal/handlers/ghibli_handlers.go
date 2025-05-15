package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"os"
	"sync"
	"time"

	"github.com/example/get-ur-ghibli/ghibli-service/internal/generation"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
)

var (
	jobStatus = make(map[string]string) // hash -> status/result
	jobMu     sync.Mutex
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

func StartKafkaConsumer() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv("KAFKA_BROKER")},
		Topic:   "ghibli-jobs",
		GroupID: "ghibli-service",
	})
	go func() {
		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				time.Sleep(time.Second)
				continue
			}
			var job map[string]string
			json.Unmarshal(m.Value, &job)
			hash := job["hash"]
			jobMu.Lock()
			jobStatus[hash] = "processing"
			jobMu.Unlock()
			// Simulate processing
			time.Sleep(2 * time.Second)
			// Mark as done with a fake result URL
			jobMu.Lock()
			jobStatus[hash] = "done:https://ghibli-service/fake-ghibli/" + hash + ".png"
			jobMu.Unlock()
		}
	}()
}

func JobStatusHandler(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query().Get("hash")
	jobMu.Lock()
	status, ok := jobStatus[hash]
	jobMu.Unlock()
	if !ok {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"` + status + `"}`))
}
