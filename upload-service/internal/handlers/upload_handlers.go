package handlers

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/segmentio/kafka-go"
)

var (
	dedupMap = make(map[string]bool)
	dedupMu  sync.Mutex
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	// Parse user from header (for demo)
	user := r.Header.Get("X-User")
	if user == "" {
		http.Error(w, "Missing user", http.StatusBadRequest)
		return
	}
	// Read image bytes
	imgBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading image", http.StatusBadRequest)
		return
	}
	// Deduplication: hash image+user
	h := sha256.New()
	h.Write(imgBytes)
	h.Write([]byte(user))
	hash := hex.EncodeToString(h.Sum(nil))
	dedupMu.Lock()
	if dedupMap[hash] {
		dedupMu.Unlock()
		http.Error(w, "Duplicate job", http.StatusConflict)
		return
	}
	dedupMap[hash] = true
	dedupMu.Unlock()
	// Upload to S3
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
	s3c := s3.New(sess)
	bucket := os.Getenv("S3_BUCKET")
	key := hash + ".png"
	_, err = s3c.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(imgBytes),
		ContentType: aws.String("image/png"),
	})
	if err != nil {
		http.Error(w, "S3 upload failed", http.StatusInternalServerError)
		return
	}
	s3url := "https://" + bucket + ".s3.amazonaws.com/" + key
	// Enqueue job to Kafka
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{os.Getenv("KAFKA_BROKER")},
		Topic:   "ghibli-jobs",
	})
	job := map[string]string{"user": user, "image_url": s3url, "hash": hash}
	jobBytes, _ := json.Marshal(job)
	err = writer.WriteMessages(r.Context(), kafka.Message{Value: jobBytes})
	if err != nil {
		http.Error(w, "Kafka enqueue failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"imageID":"%s", "s3url":"%s"}`, hash, s3url)))
}
