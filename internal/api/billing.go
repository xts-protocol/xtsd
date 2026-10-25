package api

import (
	"encoding/json"
	"net/http"

	"github.com/xts-protocol/xtsd/internal/billing"
	"github.com/xts-protocol/xtsd/internal/metrics"
	"github.com/xts-protocol/xtsd/internal/tenant"
)

func BillingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	apiKey := r.URL.Query().Get("apikey")

	t, ok := tenant.Registry[apiKey]
	if !ok {
		http.Error(w, "tenant not found", http.StatusNotFound)
		return
	}

	usage := metrics.GetUsage()

	messages := usage[apiKey]

	invoice := billing.Calculate(
		t.Name,
		messages,
	)

	json.NewEncoder(w).Encode(invoice)
}
