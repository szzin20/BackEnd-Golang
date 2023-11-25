package web

import (
	"time"
)

type ComplaintResponse struct {
	Message   string    `json:"message" form:"message"`
	Image     string    `json:"image" form:"image"`
	Audio     string    `json:"audio" form:"audio"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

type ComplaintsResponse struct {
	UserID              uint      `json:"user_id"`
	UserFullname        string    `json:"user_fullname"`
	DoctorID            uint      `json:"doctor_id"`
	DoctorFullname      string    `json:"doctor_fullname"`
	Specialist          string    `json:"specialist"`
	DoctorStatus        bool      `json:"doctor_status"`
	DoctorTransactionID uint      `json:"doctor_transaction_id"`
	CreatedAt           time.Time `json:"created_at"`
	HealthDetails       string    `json:"health_details"`
	PatientStatus       string    `json:"patient_status"`
}

type UserComplaintResponse struct {
	UserFullname        string    `json:"user_fullname"`
	Gender              string    `json:"gender"`
	Weight              int       `json:"weight"`
	CreatedAt           time.Time `json:"created_at"`
	DoctorTransactionID uint      `json:"transaction_id"`
	HealthDetails       string    `json:"health_details"`
	PatientStatus       string    `json:"patient_status"`
}
