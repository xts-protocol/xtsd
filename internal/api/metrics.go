package api

import (
	"encoding/json"
	"net/http"
	"sync/atomic"
)

var TotalMessages uint64

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]any{
		"total_messages": atomic.LoadUint64(&TotalMessages),
	})
}
