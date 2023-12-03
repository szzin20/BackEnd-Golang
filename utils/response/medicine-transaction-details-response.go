package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToCheckoutResponse(checkout *schema.Checkout) web.CheckoutResponse {

	return web.CheckoutResponse{
		ID:                    checkout.ID,
		PaymentStatus:         checkout.PaymentStatus,
		MedicineTransactionID: checkout.MedicineTransactionID,
		CreatedAt:             checkout.CreatedAt,
		PaymentConfirmation:   checkout.PaymentConfirmation,
	}
}

func ConvertToCheckoutListResponse(checkouts []schema.Checkout) []web.CheckoutResponse {
	var checkoutResponses []web.CheckoutResponse

	for _, checkout := range checkouts {
		checkoutResponses = append(checkoutResponses, ConvertToCheckoutResponse(&checkout))
	}

	return checkoutResponses
}
