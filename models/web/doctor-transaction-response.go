package web

import "time"

type CreateDoctorTransactionResponse struct {
	ID                  uint      `json:"id"  form:"id"`
	Fullname            string    `json:"fullname"  form:"fullname"`
	Specialist          string    `json:"specialist" form:"specialist"`
	Price               int       `json:"price" form:"price"`
	PaymentMethod       string    `json:"payment_method" form:"payment_method"`
	PaymentStatus       string    `json:"payment_status" form:"payment_status"`
	PaymentConfirmation string    `json:"payment_confirmation" form:"payment_confirmation"`
	CreatedAt           time.Time `json:"created_at" form:"created_at"`
}

type DoctorTransactionsResponse struct {
	ID                  uint      `json:"id"  form:"id"`
	Fullname            string    `json:"fullname"  form:"fullname"`
}
