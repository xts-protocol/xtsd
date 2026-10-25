package api

import (
	"encoding/json"
	"fmt"
	"github.com/xts-protocol/xtsd/internal/auth"
	"github.com/xts-protocol/xtsd/internal/metrics"
	"github.com/xts-protocol/xtsd/internal/receipt"
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

	messageID := fmt.Sprintf(
		"xts_%d",
		time.Now().UnixNano(),
	)

	timestamp := time.Now().
		UTC().
		Format(time.RFC3339)

	receipt.Save(
		receipt.Receipt{
			MessageID: messageID,
			APIKey:    apiKey,
			Sender:    req.Sender,
			Receiver:  req.Receiver,
			Timestamp: timestamp,
		},
	)
	response := map[string]any{
		"accepted":   true,
		"message_id": messageID,
		"sender":     req.Sender,
		"receiver":   req.Receiver,
		"timestamp":  timestamp,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
