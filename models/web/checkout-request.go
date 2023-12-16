package web

type CheckoutRequest struct {
	MedicineTransactionID uint   `json:"medicine_transaction_id" form:"medicine_transaction_id" validate:"required"`
	PaymentConfirmation   string `json:"payment_confirmation" form:"payment_confirmation"`
}

type CheckoutUpdate struct {
	PaymentStatus string `json:"payment_status" form:"payment_status" validate:"required"`
}
