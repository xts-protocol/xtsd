package api

import (
	"encoding/json"
	"net/http"

	"github.com/xts-protocol/xtsd/internal/billing"
	"github.com/xts-protocol/xtsd/internal/metrics"
	"github.com/xts-protocol/xtsd/internal/tenant"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {

	apiKey := r.URL.Query().Get("apikey")

	t, ok := tenant.Registry[apiKey]
	if !ok {
		http.Error(w, "tenant not found", http.StatusNotFound)
		return
	}

	usage := metrics.GetUsage()

	invoice := billing.Calculate(
		t.Name,
		usage[apiKey],
	)

	json.NewEncoder(w).Encode(map[string]any{
		"tenant_id":         t.ID,
		"tenant_name":       t.Name,
		"messages":          invoice.Messages,
		"amount_due":        invoice.AmountDue,
		"price_per_message": invoice.PricePerMessage,
		"xts_version":       "0.5.0",
		"status":            "running",
	})
}
