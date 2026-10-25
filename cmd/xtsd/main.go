package main

import (
	"log"
	"net/http"

	"github.com/xts-protocol/xtsd/internal/api"
)

func main() {
	http.HandleFunc("/health", api.HealthHandler)
	http.HandleFunc("/v1/send", api.SendHandler)
	http.HandleFunc("/metrics", api.MetricsHandler)
	http.HandleFunc("/v1/info", api.InfoHandler)
	http.HandleFunc("/v1/usage", api.UsageHandler)
	http.HandleFunc("/v1/billing", api.BillingHandler)
	http.HandleFunc("/v1/tenant", api.TenantHandler)
	http.HandleFunc("/v1/dashboard", api.DashboardHandler)

	log.Println("XTS Node starting on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
