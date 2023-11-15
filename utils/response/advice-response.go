package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToAdviceResponse(advice *schema.Advice) web.AdviceResponse {
	return web.AdviceResponse{
		Message:     advice.Message,
		Image:       advice.Image,
		Audio:       advice.Audio,
		CreatedAt:   advice.CreatedAt,
	}
}
