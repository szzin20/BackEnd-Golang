package web

type MedicineTransactionRequest struct {
	Name            string            `json:"name" form:"name" validate:"required"`
	Address         string            `json:"address" form:"address" validate:"required"`
	HP              string            `json:"hp" form:"hp" validate:"required"`
	PaymentMethod   string            `json:"payment_method" form:"payment_method" validate:"required"`
	MedicineDetails []MedicineDetails `json:"medicine_details" form:"medicine_details" validate:"required"`
}

type MedicineDetails struct {
	MedicineID uint `json:"medicine_id" form:"medicine_id" validate:"required"`
	Quantity   int  `json:"quantity" form:"quantity" validate:"required,min=1"`
}
