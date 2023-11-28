package web

import (
	"time"
)

type CreateAdviceResponse struct {
	TransactionID uint      `json:"transaction_id"`
	ID            uint      `json:"id"`
	Message       string    `json:"message"`
	Image         string    `json:"image"`
	Audio         string    `json:"audio"`
	CreatedAt     time.Time `json:"created_at"`
}

type AdvicesResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	Image     string    `json:"image"`
	Audio     string    `json:"audio"`
	CreatedAt time.Time `json:"created_at"`
}
