package web

type MedicineTransactionRequest struct {
	Name                   string                   `json:"name" validate:"required"`
	Address                string                   `json:"address" validate:"required"`
	HP                     string                   `json:"hp" validate:"required"`
	PaymentMethod          string                   `json:"payment_method" validate:"required"`
	MedicineDetailsRequest []MedicineDetailsRequest `json:"medicine_details" validate:"required"`
}
