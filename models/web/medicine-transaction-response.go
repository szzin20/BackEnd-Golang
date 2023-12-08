package web

import "time"

type MedicineTransactionResponse struct {
	ID                      uint                      `json:"id"`
	UserID                  uint                      `json:"user_id"`
	Name                    string                    `json:"name"`
	Address                 string                    `json:"address"`
	HP                      string                    `json:"hp"`
	PaymentMethod           string                    `json:"payment_method"`
	MedicineDetailsResponse []MedicineDetailsResponse `json:"medicine_details"`
	TotalPrice              int                       `json:"total_price"`
	StatusTransaction       string                    `json:"status_transaction"`
	CreatedAt               time.Time                 `json:"created_at"`
}

type MedicineCheckoutResponse struct {
	UserID                  uint                      `json:"user_id"`
	Name                    string                    `json:"name"`
	Address                 string                    `json:"address"`
	HP                      string                    `json:"hp"`
	PaymentMethod           string                    `json:"payment_method"`
	MedicineDetailsResponse []MedicineDetailsResponse `json:"medicine_details"`
	TotalPrice              int                       `json:"total_price"`
	StatusTransaction       string                    `json:"status_transaction"`
}

type MedicineDetailsResponse struct {
	MedicineID         uint `json:"medicine_id"`
	Quantity           int  `json:"quantity"`
	TotalPriceMedicine int  `json:"total_price_medicine"`
}
