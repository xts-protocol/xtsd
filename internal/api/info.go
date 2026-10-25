package api

import (
	"encoding/json"
	"net/http"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]any{
		"name":        "xtsd",
		"version":     "0.2.0",
		"description": "XTS Protocol Node",
		"status":      "running",
	})
}
