// @title XTS Protocol API
// @version 1.0.0-alpha
// @description Open Source Message Transport Runtime
// @host localhost:8080
// @BasePath /

package main

import (
	"github.com/xts-protocol/xtsd/internal/api"
	"github.com/xts-protocol/xtsd/internal/config"
	"github.com/xts-protocol/xtsd/internal/metrics"
	"github.com/xts-protocol/xtsd/internal/receipt"
	"github.com/xts-protocol/xtsd/internal/tenant"
	"log"
	"net/http"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
	tenant.LoadRegistry()
	metrics.Load()
	receipt.Load()
	http.HandleFunc("/health", api.HealthHandler)
	http.HandleFunc("/v1/send", api.SendHandler)
	http.HandleFunc("/metrics", api.MetricsHandler)
	http.HandleFunc("/v1/info", api.InfoHandler)
	http.HandleFunc("/v1/usage", api.UsageHandler)
	http.HandleFunc("/v1/billing", api.BillingHandler)
	http.HandleFunc("/v1/tenant", api.TenantHandler)
	http.HandleFunc("/v1/dashboard", api.DashboardHandler)
	http.HandleFunc("/v1/receipt", api.ReceiptHandler)
	http.HandleFunc(
		"/v1/apikey/create",
		api.APIKeyCreateHandler,
	)
	http.HandleFunc(
		"/v1/receipts",
		api.ReceiptsHandler,
	)
	addr := ":" + config.App.ListenPort

	log.Printf(
		"XTS Node %s starting on %s",
		config.App.NodeName,
		addr,
	)

	log.Fatal(
		http.ListenAndServe(addr, nil),
	)

}
