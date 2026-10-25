package api

import (
	"encoding/json"
	"net/http"

	"github.com/xts-protocol/xtsd/internal/metrics"
)

func UsageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(metrics.GetUsage())
}
