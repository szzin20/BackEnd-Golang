package web

import "time"

type CreateDoctorTransactionResponse struct {
	Fullname            string    `json:"fullname"  form:"fullname"`
	Specialist          string    `json:"specialist" form:"specialist"`
	Price               int       `json:"price" form:"price"`
	PaymentMethod       string    `json:"payment_method" form:"payment_method"`
	PaymentConfirmation string    `json:"payment_confirmation" form:"payment_confirmation"`
	PaymentStatus       string    `json:"payment_status" form:"payment_status"`
	CreatedAt           time.Time `json:"created_at" form:"created_at"`
}
