package api

import (
	"encoding/json"
	"net/http"

	"github.com/xts-protocol/xtsd/internal/receipt"
)

func ReceiptHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	data, ok := receipt.Get(id)
	if !ok {
		http.Error(
			w,
			"receipt not found",
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(data)
}
