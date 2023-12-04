package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToCreateComplaintMessageRequest(message web.CreateMessageRequest, RoomchatID uint, UserID uint) *schema.Message {
	return &schema.Message{
		RoomchatID: RoomchatID,
		UserID:     UserID,
		Message:    message.Message,
		Image:      message.Image,
		Audio:      message.Audio,
	}
}

func ConvertToCreateAdviceMessageRequest(message web.CreateMessageRequest, RoomchatID uint, DoctorID uint) *schema.Message {
	return &schema.Message{
		RoomchatID: RoomchatID,
		DoctorID:   DoctorID,
		Message:    message.Message,
		Image:      message.Image,
		Audio:      message.Audio,
	}
}
