package web

import "time"

type MedicineTransactionResponse struct {
	ID                      uint                      `json:"id"`
	Name                    string                    `json:"name"`
	UserID                  uint                      `json:"user_id"`
	Address                 string                    `json:"address"`
	HP                      string                    `json:"hp"`
	PaymentMethod           string                    `json:"payment_method"`
	MedicineDetailsResponse []MedicineDetailsResponse `json:"medicine_details"`
	TotalPrice              int                       `json:"total_price"`
	CreatedAt               time.Time                 `json:"created_at"`
}
