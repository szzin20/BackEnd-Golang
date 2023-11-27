package web

import (
	"time"
)

type CreateComplaintResponse struct {
	TransactionID uint      `json:"transaction_id"`
	ID            uint      `json:"id"`
	Message       string    `json:"message"`
	Image         string    `json:"image"`
	Audio         string    `json:"audio"`
	CreatedAt     time.Time `json:"created_at"`
}

type ComplaintsResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	Image     string    `json:"image"`
	Audio     string    `json:"audio"`
	CreatedAt time.Time `json:"created_at"`
}

// type ComplaintsResponse struct {
// 	UserID              uint      `json:"user_id"`
// 	UserFullname        string    `json:"user_fullname"`
// 	DoctorID            uint      `json:"doctor_id"`
// 	DoctorFullname      string    `json:"doctor_fullname"`
// 	Specialist          string    `json:"specialist"`
// 	DoctorStatus        bool      `json:"doctor_status"`
// 	DoctorTransactionID uint      `json:"doctor_transaction_id"`
// 	CreatedAt           time.Time `json:"created_at"`
// 	HealthDetails       string    `json:"health_details"`
// 	PatientStatus       string    `json:"patient_status"`
// }

// type UserComplaintResponse struct {
// 	UserFullname        string    `json:"user_fullname"`
// 	Gender              string    `json:"gender"`
// 	Weight              int       `json:"weight"`
// 	CreatedAt           time.Time `json:"created_at"`
// 	DoctorTransactionID uint      `json:"transaction_id"`
// 	HealthDetails       string    `json:"health_details"`
// 	PatientStatus       string    `json:"patient_status"`
// }
