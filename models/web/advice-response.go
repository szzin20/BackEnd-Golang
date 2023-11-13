package web

import (
	"time"
)

type AdviceResponse struct {
	ID          int       `json:"id" form:"id"`
	ComplaintID int       `json:"complaint_id" form:"complaint_id"`
	Title       string    `json:"title" form:"title"`
	Content     string    `json:"content" form:"content"`
	Status      string    `json:"status" form:"status"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
}
