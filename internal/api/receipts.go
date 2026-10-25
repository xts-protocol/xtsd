package api

import (
	"encoding/json"
	"net/http"

	"github.com/xts-protocol/xtsd/internal/receipt"
)

func ReceiptsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	apiKey := r.URL.Query().
		Get("apikey")

	items := receipt.ListByAPIKey(
		apiKey,
	)

	w.Header().
		Set(
			"Content-Type",
			"application/json",
		)

	json.NewEncoder(w).
		Encode(items)
}
