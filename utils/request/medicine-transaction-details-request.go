package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToCheckoutRequest(checkout web.CheckoutRequest) *schema.Checkout {
	return &schema.Checkout{
		MedicineTransactionID: checkout.MedicineTransactionID,
		PaymentConfirmation:   checkout.PaymentConfirmation,
	}
}

func ConvertToCheckoutUpdate(checkout web.CheckoutUpdate) *schema.Checkout {
	return &schema.Checkout{
		PaymentStatus: checkout.PaymentStatus,
	}
}
