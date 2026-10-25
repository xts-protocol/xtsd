package api

import (
	"encoding/json"
	"net/http"

	"github.com/xts-protocol/xtsd/internal/config"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(map[string]any{
		"name":        config.App.NodeName,
		"environment": config.App.Environment,
		"version":     "0.9.0",
		"status":      "running",
	})
}
