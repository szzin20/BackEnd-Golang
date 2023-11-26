package web

import "time"

type MedicineDetailsResponse struct {
	MedicineID         uint `json:"medicine_id"`
	Quantity           int  `json:"quantity"`
	TotalPriceMedicine int  `json:"total_price_medicine"`
}

type CheckoutResponse struct {
	ID                    uint      `json:"id"`
	MedicineTransactionID uint      `json:"medicine_transaction_id"`
	PaymentConfirmation   string    `json:"payment_confirmation"`
	PaymentStatus         string    `json:"payment_status"`
	CreatedAt             time.Time `json:"created_at"`
}
