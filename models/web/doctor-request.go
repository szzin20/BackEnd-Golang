package web

type DoctorRegisterRequest struct {
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
	Fullname       string `json:"fullname" form:"fullname" validate:"required"`
	Email          string `json:"email" form:"email" validate:"required,email"`
	Password       string `json:"password" form:"password" validate:"required,min=10,max=15"`
	Price          int    `json:"price" form:"price" validate:"required,min=0"`
	Gender         string `json:"gender" form:"gender" validate:"required"`
	Specialist     string `json:"specialist" form:"specialist" validate:"required"`
	NoSTR          int    `json:"no_str" form:"no_str" validate:"required"`
	Experience     string `json:"experience" form:"experience" validate:"required"`
	Alumnus        string `json:"alumnus" form:"alumnus" validate:"required"`
}

type DoctorLoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=10,max=15"`
}

type DoctorUpdateRequest struct {
	ProfilePicture string `json:"profile_picture" form:"profile_picture" validate:"omitempty"`
	Fullname       string `json:"fullname" form:"fullname" validate:"omitempty,min=3"`
	Gender         string `json:"gender" form:"gender" validate:"omitempty"`
	Email          string `json:"email" form:"email" validate:"omitempty,email"`
	Password       string `json:"password" form:"password" validate:"omitempty,min=10,max=15"`
	Price          int    `json:"price" form:"price" validate:"omitempty,min=0"`
	Specialist     string `json:"specialist" form:"specialist" validate:"omitempty"`
	Experience     string `json:"experience" form:"experience" validate:"omitempty"`
	Alumnus        string `json:"alumnus" form:"alumnus" validate:"omitempty"`
	NoSTR          int    `json:"no_str" form:"no_str" validate:"omitempty"`
	Status         bool   `json:"status" form:"status" validate:"omitempty"`
}

// Manage Patient
type UpdateManageUserRequest struct {
	HealthDetails string `json:"health_details" form:"health_details" validate:"omitempty,min=3"`
	PatientStatus string `json:"patient_status" form:"patient_status" validate:"omitempty"`
}

type ChangeDoctorStatusRequest struct {
	Status bool `json:"status" form:"status" validate:"omitempty"`
}

type PasswordResetRequest struct {
	Email string `json:"email" form:"email" validate:"required,email" `
}

type OTPVerificationRequest struct {
	OTP   string `json:"otp" `
	Email string `json:"email" form:"email" validate:"required,email"`
}

type ResetRequest struct {
	OTP      string `json:"otp" `
	Email    string `json:"email" form:"email" validate:"required,email" `
	Password string `json:"password" form:"password" validate:"omitempty,min=10,max=15"`
}
