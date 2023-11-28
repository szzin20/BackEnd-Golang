package request

import "healthcare/models/schema"

func CreateRoomchatRequest(transactionID uint) schema.Roomchat {
	return schema.Roomchat{
		TransactionID: transactionID,
	}
}
