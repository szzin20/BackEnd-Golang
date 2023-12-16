package request

import "healthcare/models/schema"

func CreateRoomchatRequest(transactionID uint) schema.Roomchat {
	return schema.Roomchat{
		TransactionID: transactionID,
	}
}

func ConvertToGetAllRoomchatRequest(userID uint,  fullname string,) *schema.DoctorTransaction {
	return &schema.DoctorTransaction{
		UserID:              userID,
	}
}