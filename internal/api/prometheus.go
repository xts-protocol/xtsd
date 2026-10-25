package api

import (
	"fmt"
	"net/http"

	"github.com/xts-protocol/xtsd/internal/metrics"
	"github.com/xts-protocol/xtsd/internal/receipt"
	"github.com/xts-protocol/xtsd/internal/tenant"
)

func PrometheusHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	usage := metrics.GetUsage()

	var messages uint64

	for _, count := range usage {
		messages += count
	}

	w.Header().Set(
		"Content-Type",
		"text/plain",
	)

	fmt.Fprintf(
		w,
		"xts_messages_total %d\n",
		messages,
	)

	fmt.Fprintf(
		w,
		"xts_tenants_total %d\n",
		len(tenant.Registry),
	)

	fmt.Fprintf(
		w,
		"xts_receipts_total %d\n",
		len(receipt.Receipts),
	)
}
