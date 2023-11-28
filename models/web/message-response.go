package web

import "time"

type CreateMessageResponse struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"user_id"`
	DoctorID   uint      `json:"doctor_id"`
	RoomchatID uint      `json:"roomchat_id"`
	Message    string    `json:"message"`
	Image      string    `json:"image"`
	Audio      string    `json:"audio"`
	CreatedAt  time.Time `json:"created_at"`
}
