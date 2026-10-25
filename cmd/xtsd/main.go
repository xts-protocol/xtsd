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

	log.Println("XTS Node starting on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
