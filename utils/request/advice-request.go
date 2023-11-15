package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToAdviceRequest(advice web.AdviceRequest, ComplaintID uint) *schema.Advice {
	return &schema.Advice{
		Message:   advice.Message,
		Image: advice.Image,
		Audio: advice.Audio,
	}
}