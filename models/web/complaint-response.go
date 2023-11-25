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
