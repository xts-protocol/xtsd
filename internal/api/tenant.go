package api

import (
	"encoding/json"
	"net/http"

	"github.com/xts-protocol/xtsd/internal/tenant"
)

func TenantHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	apiKey := r.URL.Query().Get("apikey")

	t, ok := tenant.Registry[apiKey]
	if !ok {
		http.Error(w, "tenant not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(t)
}
