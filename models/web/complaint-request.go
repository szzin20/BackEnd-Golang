package web



type ComplaintRequest struct {
	Message string `json:"message" form:"message"`
	Image   string `json:"image" form:"image"`
	Audio   string `json:"audio" form:"audio"`
}

type UpdateComplaintRequest struct {
	HealthDetails string `json:"health_details" form:"health_details" validate:"required,min=3"`
	PatientStatus string `json:"status" form:"status" validate:"required,oneof=Diberi resep Konsultasi Lanjutan Di Rujuk"`
}
