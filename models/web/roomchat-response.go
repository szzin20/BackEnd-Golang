package web

import (
	"time"
)

type CreateRoomchatResponse struct {
	ID            uint                    `json:"id"`
	TransactionID uint                    `json:"transaction_id"`
	CreatedAt     time.Time               `json:"created_at"`
}

type RoomchatDetailsResponse struct {
	ID            uint                    `json:"id"`
	TransactionID uint                    `json:"transaction_id"`
	CreatedAt     time.Time               `json:"created_at"`
	Messages      []CreateMessageResponse `json:"messages"`
}
