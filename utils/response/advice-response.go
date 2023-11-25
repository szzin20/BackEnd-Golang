package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToAdviceResponse(advice *schema.Advice) web.CreateAdviceResponse{
	return web.CreateAdviceResponse{
		TransactionID: advice.TransactionID,
		Message:       advice.Message,
		Image:         advice.Image,
		Audio:         advice.Audio,
		CreatedAt:     advice.CreatedAt,
	}
}

