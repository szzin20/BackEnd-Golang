package web



type ComplaintRequest struct {
	Message string `json:"message" form:"message"`
	Image   string `json:"image" form:"image"`
	Audio   string `json:"audio" form:"audio"`
}

type UpdateComplaintRequest struct {
	HealthDetails string `json:"health_details" form:"health_details" validate:"required,min=3"`
	PatientStatus string `json:"patient_status" form:"patient_status" validate:"required,in=pending,recovered,ongoing consultation,referred"`
}