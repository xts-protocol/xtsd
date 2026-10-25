package api

import (
	"encoding/json"
	"net/http"
)

// HealthHandler godoc
//
// @Summary Health Check
// @Description Check XTS Node status
// @Tags System
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"service": "xtsd",
	})
}
