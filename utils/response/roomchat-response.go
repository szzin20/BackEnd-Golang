package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
)

func ConvertToCreateRoomchatResponse(roomchat *schema.Roomchat) web.CreateRoomchatResponse {
	return web.CreateRoomchatResponse{
		ID:            roomchat.ID,
		TransactionID: roomchat.TransactionID,
		CreatedAt:     roomchat.CreatedAt,
	}
}

func ConvertToRoomchatUserResponse(roomchat *schema.Roomchat, doctor *schema.Doctor) web.RoomchatUserDetailsResponse {
	roomchats := web.RoomchatUserDetailsResponse{
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

	doctorProfile := web.DoctorProfileRoomchat{
		ID:             doctor.ID,
		Fullname:       doctor.Fullname,
		Status:         doctor.Status,
		ProfilePicture: doctor.ProfilePicture,
	}
	roomchats.Doctor = doctorProfile

	return roomchats
}

func ConvertToRoomchatDoctorResponse(roomchat *schema.Roomchat, user *schema.User) web.RoomchatDoctorDetailsResponse {
	roomchats := web.RoomchatDoctorDetailsResponse{
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

	userProfile := web.UserProfileRoomchat{
		ID:             user.ID,
		Fullname:       user.Fullname,
		ProfilePicture: user.ProfilePicture,
	}
	roomchats.User = userProfile

	return roomchats
}

func ConvertToGetAllRoomchats(user schema.User, roomchat schema.Roomchat, lastMessage schema.Message) web.RoomchatListResponse {
	lastMessageContent := helper.GetMessageContent(lastMessage)

	return web.RoomchatListResponse{
		ID:             roomchat.ID,
		Fullname:       user.Fullname,
		ProfilePicture: user.ProfilePicture,
		LastMessage:    lastMessageContent,
		CreatedAt:      lastMessage.CreatedAt,
	}
}
