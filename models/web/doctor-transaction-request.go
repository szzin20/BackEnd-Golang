package web

type CreateDoctorTransactionRequest struct {
	DoctorID      uint   `json:"doctor_id" form:"doctor_id"`
	UserID        uint   `json:"user_id" form:"user_id"`
	HealthDetails string `json:"health_details" form:"health_details"`
	PaymentMethod string `json:"payment_method" form:"payment_method"`
	Price         int    `json:"price" form:"price"`
	ImageURL      string `json:"image_url" form:"image_url"`
}
