package web

import "time"

type CheckoutResponse struct {
	ID                       uint                     `json:"id"`
	PaymentStatus            string                   `json:"payment_status"`
	MedicineTransactionID    uint                     `json:"medicine_transaction_id"`
	MedicineCheckoutResponse MedicineCheckoutResponse `json:"medicine_transaction"`
	CreatedAt                time.Time                `json:"created_at"`
	PaymentConfirmation      string                   `json:"payment_confirmation"`
}
