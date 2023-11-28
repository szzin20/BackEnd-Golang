package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToCreateMessageResponse(message *schema.Message) web.CreateMessageResponse {
	return web.CreateMessageResponse{
		ID:         message.ID,
		UserID:     message.UserID,
		DoctorID:   message.DoctorID,
		RoomchatID: message.RoomchatID,
		Message:    message.Message,
		Image:      message.Image,
		Audio:      message.Audio,
		CreatedAt:  message.CreatedAt,
	}
}
