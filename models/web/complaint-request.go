package web

type CreateComplaintRequest struct {
	Message string `json:"message" form:"message" validate:"omitempty"`
	Image   string `json:"image" form:"image" validate:"omitempty"`
	Audio   string `json:"audio" form:"audio" validate:"omitempty"`
}

type UpdateComplaintRequest struct {
	HealthDetails string `json:"health_details" form:"health_details" validate:"required,min=3"`
	PatientStatus string `json:"patient_status" form:"patient_status" validate:"required,in=pending,recovered,ongoing consultation,referred"`
}