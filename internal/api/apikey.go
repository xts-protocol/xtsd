package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/xts-protocol/xtsd/internal/apikey"
	"github.com/xts-protocol/xtsd/internal/tenant"
)

type CreateTenantRequest struct {
	Name string `json:"name"`
}

func APIKeyCreateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	var req CreateTenantRequest

	if err := json.NewDecoder(r.Body).
		Decode(&req); err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	key := apikey.Generate()

	t := tenant.Tenant{
		ID: fmt.Sprintf(
			"tenant-%d",
			time.Now().Unix(),
		),
		Name:   req.Name,
		APIKey: key,
	}

	tenant.Create(t)

	w.Header().
		Set(
			"Content-Type",
			"application/json",
		)

	json.NewEncoder(w).Encode(t)
}
