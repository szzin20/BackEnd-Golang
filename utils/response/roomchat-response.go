package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToCreateRoomchatResponse(roomchat *schema.Roomchat) web.CreateRoomchatResponse {
	return web.CreateRoomchatResponse{
		ID:            roomchat.ID,
		TransactionID: roomchat.TransactionID,
		CreatedAt:     roomchat.CreatedAt,
	}
}

func ConvertToRoomchatResponse(roomchat *schema.Roomchat, doctor *schema.Doctor) web.RoomchatDetailsResponse {
	roomchats := web.RoomchatDetailsResponse{
		ID:            roomchat.ID,
		TransactionID: roomchat.TransactionID,
		CreatedAt:     roomchat.CreatedAt,
	}

	var results []web.CreateMessageResponse
	for _, message := range roomchat.Message {
		roomchatResponses := web.CreateMessageResponse{
			ID:         message.ID,
			UserID:     message.UserID,
			DoctorID:   message.DoctorID,
			RoomchatID: message.RoomchatID,
			Message:    message.Message,
			Image:      message.Image,
			Audio:      message.Audio,
			CreatedAt:  message.CreatedAt,
		}
		results = append(results, roomchatResponses)
	}
	roomchats.Messages = results
	doctorprofile := web.DoctorRoomchat{
		ID:             doctor.ID,
		Fullname:       doctor.Fullname,
		Status:         doctor.Status,
		ProfilePicture: doctor.ProfilePicture,
	}
	roomchats.Doctor = doctorprofile

	return roomchats
}

func ConvertToGetAllRoomchat(roomchats []schema.Roomchat) []web.RoomchatListResponse {
	var results []web.RoomchatListResponse
	for _, roomchat := range roomchats {
		roomchatResponse := web.RoomchatListResponse{
			ID:        roomchat.ID,
			// Fullname:  user.Fullname,
			// Message:   roomchat.Message,
			CreatedAt: roomchat.CreatedAt,
		}
		results = append(results, roomchatResponse)
	}
	return results
}

func ConvertToGetRoomchats(user schema.User, roomchat schema.Roomchat, lastMessage schema.Message) web.RoomchatListResponse {
	return web.RoomchatListResponse{
		ID: roomchat.ID,
		Fullname: user.Fullname,
		LastMessage: lastMessage.Message,
		CreatedAt: lastMessage.CreatedAt,
	}
}
