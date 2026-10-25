package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

type SendRequest struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Payload  string `json:"payload"`
}

var totalMessages uint64

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"service": "xtsd",
	})
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SendRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// เพิ่มจำนวนข้อความ
	atomic.AddUint64(&totalMessages, 1)

	response := map[string]any{
		"accepted":   true,
		"message_id": fmt.Sprintf("xts_%d", time.Now().UnixNano()),
		"sender":     req.Sender,
		"receiver":   req.Receiver,
		"timestamp":  time.Now().UTC(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]any{
		"total_messages": atomic.LoadUint64(&totalMessages),
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/v1/send", sendHandler)
	http.HandleFunc("/metrics", metricsHandler)

	log.Println("XTS Node starting on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
