package web

import "time"

type DoctorRegisterResponse struct {
	ID             uint   `json:"id"`
	Fullname       string `json:"fullname"`
	Email          string `json:"email"`
	Price          int    `json:"price"`
	Gender         string `json:"gender"`
	Specialist     string `json:"specialist"`
	ProfilePicture string `json:"profile_picture"`
	NoSTR          int    `json:"no_str"`
	Experience     string `json:"experience"`
	Alumnus        string `json:"alumnus"`
}
type DoctorLoginResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type DoctorUpdateResponse struct {
	ProfilePicture string `json:"profile_picture"`
	Fullname       string `json:"fullname"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	Price          int    `json:"price"`
	Specialist     string `json:"specialist"`
	Experience     string `json:"experience"`
	Alumnus        string `json:"alumnus"`
	NoSTR          int    `json:"no_str"`
}

type DoctorAllResponse struct {
	ID             uint   `json:"id"`
	ProfilePicture string `json:"profile_picture"`
	Fullname       string `json:"fullname"`
	Specialist     string `json:"specialist"`
	Price          int    `json:"price"`
}

type DoctorAllResponseByAdmin struct {
	ID             uint   `json:"id"`
	ProfilePicture string `json:"profile_picture"`
	Fullname       string `json:"fullname"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	Price          int    `json:"price"`
	Specialist     string `json:"specialist"`
	Experience     string `json:"experience"`
	Alumnus        string `json:"alumnus"`
	NoSTR          int    `json:"no_str"`
	// DoctorTransaction []DoctorTransaction `gorm:"ForeignKey:DoctorID;references:ID"`
}
type DoctorIDResponseByAdmin struct {
	ID             uint   `json:"id" form:"id"`
	ProfilePicture string `json:"profilePicture"`
	Fullname       string `json:"fullname"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	Price          int    `json:"price"`
	Specialist     string `json:"specialist"`
	Alumnus        string `json:"alumnus"`
	Experience     string `json:"experience"`
	NoSTR          int    `json:"no_str"`
	// DoctorTransaction []DoctorTransaction `gorm:"ForeignKey:DoctorID;references:ID"`
}

type DoctorIDResponse struct {
	ID             uint   `json:"id"`
	ProfilePicture string `json:"profile_picture"`
	Status         bool   `json:"status"`
	Fullname       string `json:"fullname"`
	Specialist     string `json:"specialist"`
	Price          int    `json:"price"`
	Alumnus        string `json:"alumnus"`
	Experience     string `json:"experience"`
	NoSTR          int    `json:"no_str"`
}

// Manage Patient
type ManageUserResponse struct {
	UserID              uint      `json:"user_id"`
	Fullname            string    `json:"fullname"`
	ProfilePicture      string    `json:"profile_picture"`
	DoctorTransactionID uint      `json:"transaction_id"`
	CreatedAt           time.Time `json:"created_at"`
	HealthDetails       string    `json:"health_details"`
	PatientStatus       string    `json:"patient_status"`
}

type DoctorProfileRoomchat struct {
	ID             uint   `json:"id" form:"id"`
	Fullname       string `json:"fullname" form:"fullname"`
	Status         bool   `json:"status" form:"status"`
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
}

type DoctorConsultationResponse struct {
	UserID              uint   `json:"user_id"`
	Fullname            string `json:"fullname"`
	DoctorTransactionID uint   `json:"transaction_id"`
	Price               int    `json:"price"`
}
