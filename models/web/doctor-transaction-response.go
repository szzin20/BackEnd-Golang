package web

import "time"

type CreateDoctorTransactionResponse struct {
	ID                  uint      `json:"id"`
	Fullname            string    `json:"fullname"`
	Specialist          string    `json:"specialist"`
	Price               int       `json:"price"`
	PaymentMethod       string    `json:"payment_method"`
	PaymentStatus       string    `json:"payment_status"`
	PaymentConfirmation string    `json:"payment_confirmation"`
	CreatedAt           time.Time `json:"created_at"`
}

type DoctorTransactionsResponse struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
}

type DoctorTransactionDetailsResponse struct {
	ID            uint   `json:"id"`
	Fullname      string `json:"fullname"`
	Specialist    string `json:"specialist"`
	PatientStatus string `json:"patient_status"`
}

type GetUserTransactionbyAdminResponse struct {
	ID                  uint      `json:"id"`
	DoctorID            uint      `json:"doctor_id"`
	PaymentMethod       string    `json:"payment_method"`
	Price               int       `json:"price"`
	CreatedAt           time.Time `json:"created_at"`
	PaymentConfirmation string    `json:"payment_confirmation"`
	PaymentStatus       string    `json:"payment_status"`
}
