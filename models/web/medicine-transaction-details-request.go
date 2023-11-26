package web

type MedicineDetailsRequest struct {
	MedicineID uint `json:"medicine_id" form:"medicine_id" validate:"required"`
	Quantity   int  `json:"quantity" form:"quantity" validate:"required"`
}

type CheckoutRequest struct {
	MedicineTransactionID uint   `json:"medicine_transaction_id" form:"medicine_transaction_id" validate:"required"`
	PaymentConfirmation   string `json:"payment_confirmation" form:"payment_confirmation" validate:"required"`
}
