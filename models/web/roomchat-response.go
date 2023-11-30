package web

import (
	"time"
)

type CreateRoomchatResponse struct {
	ID            uint      `json:"id"`
	TransactionID uint      `json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
}

type RoomchatListResponse struct {
	ID          uint      `json:"id"`
	Fullname    string    `json:"fullname"`
	LastMessage string    `json:"last_message"`
	CreatedAt   time.Time `json:"created_at"`
}

type RoomchatDetailsResponse struct {
	Doctor        DoctorRoomchat          `json:"doctor"`
	ID            uint                    `json:"id"`
	TransactionID uint                    `json:"transaction_id"`
	CreatedAt     time.Time               `json:"created_at"`
	Messages      []CreateMessageResponse `json:"messages"`
}
