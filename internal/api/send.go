package api

import (
	"encoding/json"
	"fmt"
	"github.com/xts-protocol/xtsd/internal/auth"
	"github.com/xts-protocol/xtsd/internal/metrics"
	"net/http"
	"sync/atomic"
	"time"
)

type SendRequest struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Payload  string `json:"payload"`
}

func SendHandler(w http.ResponseWriter, r *http.Request) {

	apiKey := r.Header.Get("X-API-Key")

	if !auth.ValidateAPIKey(apiKey) {
		http.Error(w, "invalid api key", http.StatusUnauthorized)
		return
	}
	metrics.Increment(apiKey)
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SendRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	atomic.AddUint64(&TotalMessages, 1)

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
