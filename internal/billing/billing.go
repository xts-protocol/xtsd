package billing

import "math"

const PricePerMessage = 0.000001

type Invoice struct {
	Tenant          string  `json:"tenant"`
	Messages        uint64  `json:"messages"`
	PricePerMessage float64 `json:"price_per_message"`
	AmountDue       float64 `json:"amount_due"`
}

func Calculate(tenant string, messages uint64) Invoice {

	amount := float64(messages) * PricePerMessage

	// ปัดทศนิยม 12 ตำแหน่ง
	amount = math.Round(amount*1e12) / 1e12

	return Invoice{
		Tenant:          tenant,
		Messages:        messages,
		PricePerMessage: PricePerMessage,
		AmountDue:       amount,
	}
}
