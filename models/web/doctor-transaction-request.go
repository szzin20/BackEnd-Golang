package web

type CreateDoctorTransactionRequest struct {
	PaymentMethod       string `json:"payment_method" form:"payment_method"`
	PaymentConfirmation string `json:"payment_confirmation" form:"payment_confirmation"`
}
