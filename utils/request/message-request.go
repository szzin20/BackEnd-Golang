package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToCreateMessageRequest(message web.CreateMessageRequest, RoomchatID uint) *schema.Message {
	return &schema.Message{
		RoomchatID: RoomchatID,
		Message:    message.Message,
		Image:      message.Image,
		Audio:      message.Audio,
	}
}
