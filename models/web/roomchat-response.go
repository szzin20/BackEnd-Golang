package web

import (
	"time"
)

type CreateRoomchatResponse struct {
	ID             uint       `json:"id"`
	TransactionID  uint       `json:"transaction_id"`
	Status         bool       `json:"status"`
	ExpirationTime *time.Time `json:"expiration_time"`
	CreatedAt      time.Time  `json:"created_at"`
}

type RoomchatListResponse struct {
	ID             uint       `json:"id"`
	Fullname       string     `json:"fullname"`
	ProfilePicture string     `json:"profile_picture"`
	LastMessage    string     `json:"last_message"`
	Status         bool       `json:"status"`
	ExpirationTime *time.Time `json:"expiration_time"`
	CreatedAt      time.Time  `json:"created_at"`
}

type RoomchatUserDetailsResponse struct {
	Doctor         DoctorProfileRoomchat   `json:"doctor"`
	ID             uint                    `json:"id"`
	TransactionID  uint                    `json:"transaction_id"`
	CreatedAt      time.Time               `json:"created_at"`
	Status         bool                    `json:"status"`
	ExpirationTime *time.Time              `json:"expiration_time"`
	Messages       []CreateMessageResponse `json:"messages"`
}

type RoomchatDoctorDetailsResponse struct {
	User           UserProfileRoomchat     `json:"user"`
	ID             uint                    `json:"id"`
	TransactionID  uint                    `json:"transaction_id"`
	CreatedAt      time.Time               `json:"created_at"`
	Status         bool                    `json:"status"`
	ExpirationTime *time.Time              `json:"expiration_time"`
	Messages       []CreateMessageResponse `json:"messages"`
}
