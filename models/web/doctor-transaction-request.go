package web

type CreateDoctorTransactionRequest struct {
	PaymentMethod       string `json:"payment_method" form:"payment_method" validate:"required"`
	PaymentConfirmation string `json:"payment_confirmation" form:"payment_confirmation" validate:"required"`
}
