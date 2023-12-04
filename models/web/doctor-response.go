package web

import "time"

type DoctorRegisterResponse struct {
	ID             uint   `json:"id"`
	Fullname       string `json:"fullname"`
	Email          string `json:"email"`
	Status         bool   `json:"status"`
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
	Fullname         string `json:"fullname"`
	Email            string `json:"email"`
	Gender           string `json:"gender"`
	Specialist       string `json:"specialist"`
	ProfilePicture   string `json:"profile_picture"`
	NoSTR            int    `json:"no_str"`
	Experience       string `json:"experience"`
	Alumnus          string `json:"alumnus"`
	Status           bool   `json:"status"`
	AboutDoctor      string `json:"about_doctor"`
	LocationPractice string `json:"location_practice"`
}

type DoctorAllResponse struct {
	ID             uint   `json:"id"`
	ProfilePicture string `json:"profile_picture"`
	Fullname       string `json:"fullname"`
	Specialist     string `json:"specialist"`
	Price          int    `json:"price"`
	Status         bool   `json:"status"`
}

type DoctorAllResponseByAdmin struct {
	ID             uint   `json:"id"`
	ProfilePicture string `json:"profile_picture"`
	Fullname       string `json:"fullname"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	Status         bool   `json:"status"`
	Price          int    `json:"price"`
	Specialist     string `json:"specialist"`
	Experience     string `json:"experience"`
	NoSTR          int    `json:"no_str"`
	Alumnus        string `json:"alumnus"`
	// DoctorTransaction []DoctorTransaction `gorm:"ForeignKey:DoctorID;references:ID"`
}

type DoctorIDResponse struct {
	ID               uint   `json:"id"`
	ProfilePicture   string `json:"profile_picture"`
	Status           bool   `json:"status"`
	Fullname         string `json:"fullname"`
	Specialist       string `json:"specialist"`
	Price            int    `json:"price"`
	Experience       string `json:"experience"`
	AboutDoctor      string `json:"about_doctor"`
	NoSTR            int    `json:"no_str"`
	LocationPractice string `json:"location_practice"`
	Alumnus          string `json:"alumnus"`
}

// Manage Patient
type ManageUserResponse struct {
	UserID              uint      `json:"user_id"`
	Fullname            string    `json:"fullname"`
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
