package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToCheckoutResponse(checkout *schema.Checkout) web.CheckoutResponse {
	return web.CheckoutResponse{
		ID:                    checkout.ID,
		MedicineTransactionID: checkout.MedicineTransactionID,
		PaymentConfirmation:   checkout.PaymentConfirmation,
		PaymentStatus:         checkout.PaymentStatus,
		CreatedAt:             checkout.CreatedAt,
	}
}
