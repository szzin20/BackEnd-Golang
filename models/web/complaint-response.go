package web

import (
	"time"
)

type ComplaintResponse struct {
	ID            int       `json:"id" form:"id"`
	TransactionID int       `json:"transaction_id" form:"transaction_id"`
	Title         string    `json:"title" form:"title"`
	Content       string    `json:"content" form:"content"`
	Status        string    `json:"status" form:"status"`
	CreatedAt     time.Time `json:"created_at" form:"created_at"`
}
