package web

import "time"

type AdminLoginResponse struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}
type AdminUpdateResponse struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

type AdminProfileResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AdminTransactionUsersResponse struct {
	TransactionID       uint      `json:"transaction_id"`
	DoctorID            uint      `json:"Doctor_id"`
	UserID              uint      `json:"user_id"`
	PaymentMethod       string    `json:"payment_method"`
	Price               int       `json:"price"`
	CreatedAt           time.Time `json:"created_at"`
	PaymentConfirmation string    `json:"payment_confirmation"`
	PaymentStatus       string    `json:"payment_status"`
}
