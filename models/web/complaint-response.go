package web

import (
	"time"
)

type ComplaintResponse struct {
	Message       string    `json:"message" form:"message"`
	Image         string    `json:"image" form:"image"`
	Audio         string    `json:"audio" form:"audio"`
	CreatedAt     time.Time `json:"created_at" form:"created_at"`
}
